-- name: ListMessages :many
SELECT *
  FROM messages
  WHERE id = $1;
