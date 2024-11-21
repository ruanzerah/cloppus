-- name: ListMessages :many
SELECT m.*
  FROM messages m
  JOIN users u ON m.id = u.id
  WHERE u.id = $1;
