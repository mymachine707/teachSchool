-- name: CreateEnteries :one
INSERT INTO enteries (
    accaunts_id,
    amount
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetEnteries :one
SELECT * FROM enteries
WHERE id =$1 LIMIT 1;

-- name: ListEnteriess :many
SELECT * FROM enteries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEnteries :one
UPDATE enteries
SET amount= $2
WHERE id= $1
RETURNING *;

-- name: DeleteEnteries :exec
DELETE FROM enteries WHERE id =$1;