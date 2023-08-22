package ingestion

import (
	"challenge/pkg/accounts"
	mockdb "challenge/pkg/db/mock"
	"challenge/pkg/utils"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type MockDelayGenerator struct {
	Delay time.Duration
}

func (m *MockDelayGenerator) GenerateDelay() time.Duration {
	return m.Delay
}

func (m *MockDelayGenerator) SetDelay(delay time.Duration) {
	m.Delay = delay
}

func TestProcessAccountUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockdb.NewMockRepo(ctrl)
	logger := utils.NewTestLogger()
	ctx := utils.ContextWithTestLogger(context.Background(), logger)

	update := &accounts.AccountUpdate{
		ID:             "testID",
		AccountType:    "testAccount",
		Tokens:         100,
		CallbackTimeMs: 50,
		Data:           map[string]interface{}{"testKey": "testValue"},
		Version:        1,
	}

	// Define what the mock repository should expect
	repo.EXPECT().GetAccountVersion(gomock.Any(), gomock.Eq(update.ID)).Return(sql.NullInt32{Int32: 1, Valid: true}, nil)
	repo.EXPECT().UpsertAccountUpdate(gomock.Any(), gomock.Any()).Return(nil)

	// Call the function to be tested
	err := processAccountUpdate(ctx, update, repo)

	// Check that there was no error and the expected log message was printed
	assert.NoError(t, err)
	expectedLog := "AccountId: testID, Version: 1 has been ingested"
	fmt.Println(logger.Logs())
	assert.Contains(t, logger.Logs(), expectedLog)
}
