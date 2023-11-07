-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  from_account_owner,
  to_account_id,
  to_account_owner,
  amount
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
WHERE 
    from_account_owner = $1 OR
    to_account_owner = $2
ORDER BY id
LIMIT $3
OFFSET $4;