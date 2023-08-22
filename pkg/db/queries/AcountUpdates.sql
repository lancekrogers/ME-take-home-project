-- name: InsertAccountUpdate :exec
INSERT INTO account_updates (id, account_type, tokens, data, version)
VALUES ($1, $2, $3, $4, $5);

-- name: GetAllUpdatesForAccount :many
SELECT * FROM account_updates WHERE id = $1 ORDER BY version;
