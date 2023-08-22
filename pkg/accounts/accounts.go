package accounts

import (
	"challenge/pkg/db"
	"challenge/pkg/utils"
	"context"
	"database/sql"
	"log"
)

type AccountUpdate struct {
	ID             string      `json:"id"`
	AccountType    string      `json:"accountType"`
	Tokens         int         `json:"tokens"`
	CallbackTimeMs int         `json:"callbackTimeMs"`
	Data           interface{} `json:"data"`
	Version        int64       `json:"version"`
}

type AccountResult struct {
	ID          string                 `json:"id"`
	AccountType string                 `json:"accountType"`
	Tokens      int                    `json:"tokens"`
	Data        map[string]interface{} `json:"data"`
	Version     int64                  `json:"version"`
}

func GetCurrentAccountVersion(ctx context.Context, id string, repo db.Repo) (versionNumber int64, found bool, rr error) {
	currentVersion, err := repo.GetAccountVersion(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, false, nil
		}
		return 0, false, err
	}
	return int64(currentVersion.Int32), true, nil
}

func GetAccountString(ctx context.Context, accountID string, repo db.Repo) (*AccountResult, error) {
	logger, ok := utils.LoggerFromContext(ctx)
	if !ok {
		logger = &log.Logger{}
	}

	account, err := repo.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &AccountResult{}, err
		}
	}

	result := &AccountResult{
		ID:          account.ID,
		AccountType: account.AccountType,
		Tokens:      int(account.Tokens),
		Version:     int64(account.Version.Int32),
	}

	data, err := utils.DecodeJson(account.Data.RawMessage)
	if err != nil {
		logger.Printf("Data could not be decoded for account: %s", accountID)
	} else {
		result.Data = data
	}
	return result, nil
}

func LogRichestAccounts(ctx context.Context, repo db.Repo) {
	logger, ok := utils.LoggerFromContext(ctx)
	if !ok {
		logger = &log.Logger{}
	}

	richestAccounts, err := repo.GetRichestAccountsByAccountType(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, account := range richestAccounts {
		logger.Printf("Richest %s account: Tokens: %d, ID: %v", account.AccountType, account.Tokens, account.ID)
	}
}
