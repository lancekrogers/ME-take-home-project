// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: AcountUpdates.sql

package db

import (
	"context"
	"database/sql"

	"github.com/sqlc-dev/pqtype"
)

const getAllUpdatesForAccount = `-- name: GetAllUpdatesForAccount :many
SELECT id, account_type, tokens, data, version, created_at FROM account_updates WHERE id = $1 ORDER BY version
`

func (q *Queries) GetAllUpdatesForAccount(ctx context.Context, id string) ([]AccountUpdates, error) {
	rows, err := q.db.QueryContext(ctx, getAllUpdatesForAccount, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountUpdates{}
	for rows.Next() {
		var i AccountUpdates
		if err := rows.Scan(
			&i.ID,
			&i.AccountType,
			&i.Tokens,
			&i.Data,
			&i.Version,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAccountUpdate = `-- name: InsertAccountUpdate :exec
INSERT INTO account_updates (id, account_type, tokens, data, version)
VALUES ($1, $2, $3, $4, $5)
`

type InsertAccountUpdateParams struct {
	ID          string                `json:"id"`
	AccountType string                `json:"account_type"`
	Tokens      int64                 `json:"tokens"`
	Data        pqtype.NullRawMessage `json:"data"`
	Version     sql.NullInt32         `json:"version"`
}

func (q *Queries) InsertAccountUpdate(ctx context.Context, arg InsertAccountUpdateParams) error {
	_, err := q.db.ExecContext(ctx, insertAccountUpdate,
		arg.ID,
		arg.AccountType,
		arg.Tokens,
		arg.Data,
		arg.Version,
	)
	return err
}