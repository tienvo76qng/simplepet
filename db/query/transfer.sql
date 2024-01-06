-- name: CreateTranfer :one
insert into transfers (
  from_account_id,
  to_account_id,
  amount
) values (
  $1, $2, $3
) returning *;

-- name: GetTranfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTranfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTranfer :one
UPDATE transfers
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTranfer :exec
DELETE FROM transfers
WHERE id = $1;