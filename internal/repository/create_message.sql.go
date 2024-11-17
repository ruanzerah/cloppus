// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: create_message.sql

package repository

import (
	"context"
)

const createMessage = `-- name: CreateMessage :exec
INSERT INTO messages (
  owner, subject, content
) VALUES($1, $2, $3)
`

type CreateMessageParams struct {
	Owner   string `json:"owner"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) error {
	_, err := q.db.Exec(ctx, createMessage, arg.Owner, arg.Subject, arg.Content)
	return err
}