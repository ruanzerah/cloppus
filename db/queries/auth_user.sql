-- name: AuthenticateUser :exec
UPDATE users
  SET auth = true
  WHERE id = $1;
