// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: list_messages.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const listMessages = `-- name: ListMessages :many
SELECT m.id, m.owner, m.subject, m.content, m.likes, m.created_at, m.updated_at
  FROM messages m
  JOIN users u ON m.id = u.id
  WHERE u.id = $1
`

func (q *Queries) ListMessages(ctx context.Context, id uuid.UUID) ([]Message, error) {
	rows, err := q.db.Query(ctx, listMessages, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Subject,
			&i.Content,
			&i.Likes,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
