// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	GetAccount(ctx context.Context, id string) (Accounts, error)
	GetAccountVersion(ctx context.Context, id string) (sql.NullInt32, error)
	GetAllUpdatesForAccount(ctx context.Context, id string) ([]AccountUpdates, error)
	GetRichestAccountsByAccountType(ctx context.Context) ([]GetRichestAccountsByAccountTypeRow, error)
	InsertAccountUpdate(ctx context.Context, arg InsertAccountUpdateParams) error
	UpsertAccount(ctx context.Context, arg UpsertAccountParams) error
}

var _ Querier = (*Queries)(nil)