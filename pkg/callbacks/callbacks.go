package callbacks

import (
	"challenge/pkg/accounts"
	"challenge/pkg/db"
	"challenge/pkg/utils"

	"context"
	"log"
	"sync"
	"time"
)

var (
	mu     sync.Mutex
	timers = make(map[string]*time.Timer)
)

func Schedule(ctx context.Context, accountID string, version int64, callbackTimeMs int, repo db.Repo) {
	ctx, cancel := context.WithCancel(ctx)

	logger, ok := utils.LoggerFromContext(ctx)
	if !ok {
		logger = &log.Logger{}
	}

	mu.Lock()
	if timer, exists := timers[accountID]; exists {
		if timer.Stop() {
			logger.Printf("Existing callback for accountID %s has been cancelled\n", accountID)
		}
	}
	timer := time.NewTimer(time.Duration(callbackTimeMs) * time.Millisecond)
	timers[accountID] = timer
	mu.Unlock()

	go func() {
		<-timer.C
		accountState, err := accounts.GetAccountString(ctx, accountID, repo)
		if err != nil {
			cancel()
		}
		logger.Printf("Callback log: %+v\n", accountState)
	}()

}
