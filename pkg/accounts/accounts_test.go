package accounts

import (
	mockdb "challenge/pkg/db/mock"
	db "challenge/pkg/db/sqlc"
	"challenge/pkg/utils"
	"context"
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/sqlc-dev/pqtype"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetCurrentAccountVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockdb.NewMockRepo(ctrl)
	accountID := "testID"
	version := int64(42)

	repo.EXPECT().GetAccountVersion(gomock.Any(), accountID).Return(sql.NullInt32{Int32: int32(version), Valid: true}, nil)

	ctx := context.Background()
	gotVersion, found, err := GetCurrentAccountVersion(ctx, accountID, repo)

	assert.Equal(t, version, gotVersion)
	assert.True(t, found)
	assert.NoError(t, err)
}

func TestGetAccountString(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockdb.NewMockRepo(ctrl)
	accountID := "testID"

	rawJSON := `{"testField1": "images"}`
	account := db.Accounts{
		ID:          accountID,
		AccountType: "testType",
		Tokens:      100,
		Data:        pqtype.NullRawMessage{RawMessage: json.RawMessage(rawJSON), Valid: true},
		Version:     sql.NullInt32{Int32: 1, Valid: true},
	}

	decodedData := map[string]interface{}{
		"testField1": "images",
	}

	repo.EXPECT().GetAccount(gomock.Any(), accountID).Return(account, nil)
	ctx := utils.ContextWithTestLogger(context.Background(), utils.NewTestLogger())
	gotAccount, err := GetAccountString(ctx, accountID, repo)

	assert.Equal(t, decodedData, gotAccount.Data)
	assert.NoError(t, err)
}
