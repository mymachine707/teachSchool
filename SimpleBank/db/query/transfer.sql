-- name: CreateTransfers :one
INSERT INTO transfers (
    from_accaunts_id,
    to_accaunts_id,
    amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTransfers :one
SELECT * FROM transfers
WHERE id =$1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
WHERE
    from_accaunts_id = $1 OR
    to_accaunts_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;
