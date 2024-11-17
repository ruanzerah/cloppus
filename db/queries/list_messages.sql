-- name: ListMessages :many
SELECT *
  FROM messages
  WHERE owner = $1;
