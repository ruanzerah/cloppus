-- name: CreateMessage :exec
INSERT INTO messages (
  owner, subject, content
) VALUES($1, $2, $3);
