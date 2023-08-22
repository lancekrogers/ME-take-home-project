package ingestion

import (
	"challenge/pkg/accounts"
	"challenge/pkg/callbacks"
	"challenge/pkg/db"
	"challenge/pkg/utils"
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

type DelayGenerator interface {
	GenerateDelay() time.Duration
}

type RandomDelayGenerator struct{}

func (r *RandomDelayGenerator) GenerateDelay() time.Duration {
	return time.Duration(rand.Intn(1000)) * time.Millisecond
}

func Start(ctx context.Context, delay DelayGenerator, input string, repo db.Repo) {
	logger, ok := utils.LoggerFromContext(ctx)
	if !ok {
		logger = &log.Logger{}
	}

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var accountUpdates []accounts.AccountUpdate
	if err := json.NewDecoder(file).Decode(&accountUpdates); err != nil {
		log.Fatal(err)
	}

	updatesChannel := make(chan accounts.AccountUpdate, len(accountUpdates))

	go func() {
		for update := range updatesChannel {
			err := processAccountUpdate(ctx, &update, repo)
			if err != nil {
				logger.Printf("Error processing update: %v", err)
			}
		}
	}()

	for _, update := range accountUpdates {
		time.Sleep(delay.GenerateDelay())
		updatesChannel <- update
	}

	accounts.LogRichestAccounts(ctx, repo)

	close(updatesChannel)
}

func processAccountUpdate(ctx context.Context, update *accounts.AccountUpdate, repo db.Repo) error {
	logger, ok := utils.LoggerFromContext(ctx)
	if !ok {
		logger = &log.Logger{}
	}

	// Get current account version
	currentVersion, found, err := accounts.GetCurrentAccountVersion(ctx, update.ID, repo)
	if err != nil {
		return err
	}
	if found && update.Version <= currentVersion {
		// Ignore update if the current account version in the db is newer than the one in the update
		return nil
	}

	data, err := utils.EncodeStructForDB(update.Data)
	if err != nil {
		logger.Printf("Failed to prepare data field for db, Error: %v", err)
	}
	err = repo.UpsertAccountUpdate(ctx, &db.UpsertActUpdateParams{
		ID:          update.ID,
		AccountType: update.AccountType,
		Tokens:      int64(update.Tokens),
		Data:        data,
		Version:     int32(update.Version),
	})
	if err != nil {
		return err
	}

	logger.Printf("AccountId: %s, Version: %v has been ingested", update.ID, update.Version)

	callbacks.Schedule(ctx, update.ID, update.Version, update.CallbackTimeMs, repo)
	return nil
}
