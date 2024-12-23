// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: delete_message.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const deleteMessage = `-- name: DeleteMessage :exec
DELETE FROM messages
  WHERE id = $1
`

func (q *Queries) DeleteMessage(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteMessage, id)
	return err
}
