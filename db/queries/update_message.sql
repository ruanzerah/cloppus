-- name: UpdateMessage :one
UPDATE messages
  SET subject = $3,
  content = $4
  WHERE id = $1 AND owner = $2
  RETURNING *;


