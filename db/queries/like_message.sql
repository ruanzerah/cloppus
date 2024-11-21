-- name: LikeMessage :exec
UPDATE messages
  SET likes = likes + 1
  WHERE id = $1;
