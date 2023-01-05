-- name: CreateAccount :one
INSERT INTO accaunts (
    owner,
    balance,
    currency
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accaunts
WHERE id =$1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accaunts
WHERE id =$1 LIMIT 1
FOR NO KEY UPDATE;


-- name: ListAccounts :many
SELECT * FROM accaunts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accaunts
SET balance= $2
WHERE id= $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE accaunts
SET balance= balance + sqlc.arg(amount)
WHERE id= sqlc.arg(id)
RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM accaunts WHERE id =$1;