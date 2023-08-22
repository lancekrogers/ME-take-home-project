package pkg

import (
	"challenge/pkg/db"
	"challenge/pkg/ingestion"
	"challenge/pkg/utils"
	"context"
	"embed"
	"log"

	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type Config struct {
	DatabaseSettings db.DatabaseSettings
}

type TrxIndexingService struct {
	repo db.Repo
}

//go:embed db/migrations/*.sql
var fs embed.FS

func NewTrxIndexingService(ctx context.Context, cfg *Config) (*TrxIndexingService, error) {
	logger, ok := utils.LoggerFromContext(ctx)
	if !ok {
		logger = &log.Logger{}
	}

	d, err := iofs.New(fs, "db/migrations")
	if err != nil {
		logger.Fatal("New IOFS Error", "err", err.Error())
	}

	database, err := db.MigrateAndGetDatabaseWithIOFS(d, cfg.DatabaseSettings)
	if err != nil {
		logger.Fatal("Error Migrating and Getting Database", "err", err.Error())
	}

	repo := db.NewRepo(database)

	return &TrxIndexingService{repo: repo}, nil
}

func (trx *TrxIndexingService) Run(ctx context.Context) {
	generator := &ingestion.RandomDelayGenerator{}
	ingestion.Start(ctx, generator, "data/challenge-input.json", trx.repo)
}
