-- name: CreateUser :exec
INSERT INTO users (
  username, email, auth , hash
) VALUES ( $1, $2 ,$3 ,$4 );
