-- name: CreateLoanTransfer :one
INSERT INTO loan_transfers (
  account_id,
  loan_amount,
  interest_rate,
  status,
  end_date
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetLoanTransfer :one
SELECT * FROM loan_transfers
WHERE id = $1 LIMIT 1;