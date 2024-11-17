-- name: DeleteMessage :exec
DELETE FROM messages
  WHERE id = $1;
