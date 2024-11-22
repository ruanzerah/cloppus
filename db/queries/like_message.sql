-- name: LikeMessage :exec
UPDATE messages
  SET likes = likes + 1,
  updated_at = CURRENT_TIMESTAMP
  WHERE id = $1;
