// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: list_user.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const listUserById = `-- name: ListUserById :one
SELECT id, username, email, created_at, updated_at
  FROM users
  WHERE id = $1
`

func (q *Queries) ListUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, listUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUserByName = `-- name: ListUserByName :one
SELECT id, username, email, created_at, updated_at
  FROM users
  WHERE username = $1
`

func (q *Queries) ListUserByName(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, listUserByName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
