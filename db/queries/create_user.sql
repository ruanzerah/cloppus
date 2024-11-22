-- name: CreateUser :exec
INSERT INTO users (
  username, email, created_at, updated_at
) VALUES ( $1, $2, $3, $4 );
