// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: payment.sql

package db

import (
	"context"
	"database/sql"
)

const createPayment = `-- name: CreatePayment :execresult
INSERT INTO payment (id, name, amount, event_id, member_id) VALUES (?, ?, ?, ?, ?)
`

type CreatePaymentParams struct {
	ID       string
	Name     string
	Amount   int64
	EventID  string
	MemberID string
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createPayment,
		arg.ID,
		arg.Name,
		arg.Amount,
		arg.EventID,
		arg.MemberID,
	)
}
