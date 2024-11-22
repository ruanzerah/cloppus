-- name: CreateMessage :exec
INSERT INTO messages (
  owner, subject, content, created_at, updated_at
) VALUES($1, $2, $3, $4, $5 );
