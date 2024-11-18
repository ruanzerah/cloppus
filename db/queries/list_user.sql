-- name: ListUser :one
SELECT *
  FROM users
  WHERE id = $1;
