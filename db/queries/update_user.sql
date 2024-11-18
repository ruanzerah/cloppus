-- name: RenameUser :one
UPDATE users
 SET username = $2
WHERE id = $1
RETURNING *;

-- name: ChangePassword :one
UPDATE users
  SET hash = $2
WHERE id = $1
RETURNING *;
