// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: member.sql

package db

import (
	"context"
	"database/sql"
)

const createMember = `-- name: CreateMember :execresult
INSERT INTO member (id, name, event_id) VALUES (?, ?, ?)
`

type CreateMemberParams struct {
	ID      string
	Name    string
	EventID string
}

func (q *Queries) CreateMember(ctx context.Context, arg CreateMemberParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createMember, arg.ID, arg.Name, arg.EventID)
}