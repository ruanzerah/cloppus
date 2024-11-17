-- name: UpdateMessage :exec
UPDATE messages
  SET subject = $2,
  content = $3
  WHERE id = $1;


