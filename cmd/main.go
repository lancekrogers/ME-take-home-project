package main

import (
	"challenge/pkg"
	"challenge/pkg/utils"
	"context"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig[pkg.Config]("./cmd/")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)
	// logger.SetPrefix("2023-01-01 00:00:00.000 ")

	ctx, cancel := context.WithCancel(context.Background())

	ctx = utils.ContextWithLogger(ctx, logger)

	service, err := pkg.NewTrxIndexingService(ctx, config)
	if err != nil {
		log.Fatal("failed to start indexing service", "error", err)
	}

	service.Run(ctx)

	cancel()
}
