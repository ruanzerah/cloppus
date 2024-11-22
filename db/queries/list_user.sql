-- name: ListUserById :one
SELECT *
  FROM users
  WHERE id = $1;

-- name: ListUserByName :one
SELECT *
  FROM users
  WHERE username = $1;
