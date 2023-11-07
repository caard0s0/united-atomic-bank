-- name: CreateLoanTransfer :one
INSERT INTO loan_transfers (
  account_id,
  amount
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetLoanTransfer :one
SELECT * FROM loan_transfers
WHERE id = $1 LIMIT 1;