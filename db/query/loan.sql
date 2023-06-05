-- name: CreateLoan :one
INSERT INTO loans (
  account_id,
  loan_amount,
  interest_rate,
  status,
  end_date
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetLoan :one
SELECT * FROM loans
WHERE id = $1 LIMIT 1;