-- name: UpdateMessage :one
UPDATE messages
  SET subject = $3,
  content = $4,
  updated_at = CURRENT_TIMESTAMP
  WHERE id = $1 AND owner = $2
  RETURNING *;


