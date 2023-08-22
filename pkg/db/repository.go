package db

import (
	db "challenge/pkg/db/sqlc"
	"context"
	"database/sql"
	"fmt"

	"github.com/sqlc-dev/pqtype"
)

type Repo interface {
	db.Querier
	UpsertAccountUpdate(ctx context.Context, arg *UpsertActUpdateParams) error
}

type SQLRepo struct {
	*db.Queries
	db *sql.DB
}

func NewRepo(database *sql.DB) *SQLRepo {
	return &SQLRepo{
		Queries: db.New(database),
		db:      database,
	}
}

func (repo *SQLRepo) execTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type UpsertActUpdateParams struct {
	ID          string                `json:"id"`
	AccountType string                `json:"account_type"`
	Tokens      int64                 `json:"tokens"`
	Data        pqtype.NullRawMessage `json:"data"`
	Version     int32                 `json:"version"`
}

func (repo *SQLRepo) UpsertAccountUpdate(ctx context.Context, arg *UpsertActUpdateParams) error {
	version := sql.NullInt32{
		Int32: arg.Version,
		Valid: true,
	}
	// Insert update into AccountUpdates table
	err := repo.execTx(ctx, func(q *db.Queries) error {
		err := q.InsertAccountUpdate(ctx, db.InsertAccountUpdateParams{
			ID:          arg.ID,
			AccountType: arg.AccountType,
			Tokens:      arg.Tokens,
			Data:        arg.Data,
			Version:     version,
		})
		return err
	})
	if err != nil {
		fmt.Printf("AccountUpdates insert error %v", err)
		return err
	}
	// Upsert update into accounts table
	err = repo.execTx(ctx, func(q *db.Queries) error {
		err := q.UpsertAccount(ctx, db.UpsertAccountParams{
			ID:          arg.ID,
			AccountType: arg.AccountType,
			Tokens:      arg.Tokens,
			Data:        arg.Data,
			Version:     version,
		})
		return err
	})
	return err
}
