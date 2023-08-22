-- name: GetRichestAccountsByAccountType :many
SELECT account_type, tokens, id 
FROM accounts
GROUP BY account_type, tokens, id
HAVING tokens = MAX(tokens);

-- name: UpsertAccount :exec
INSERT INTO accounts (id, account_type, tokens, data, version)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (id)
DO UPDATE SET account_type=$2, tokens=$3, data=$4, version=$5, updated_at=now()
WHERE accounts.version <= $5;


-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1;

-- name: GetAccountVersion :one
SELECT version FROM accounts
WHERE id = $1;
