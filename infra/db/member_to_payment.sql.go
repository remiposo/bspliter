// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: member_to_payment.sql

package db

import (
	"context"
	"database/sql"
)

const createMemberToPayment = `-- name: CreateMemberToPayment :execresult
INSERT INTO member_to_payment (event_id, member_id, payment_id) VALUES (?, ?, ?)
`

type CreateMemberToPaymentParams struct {
	EventID   string
	MemberID  string
	PaymentID string
}

func (q *Queries) CreateMemberToPayment(ctx context.Context, arg CreateMemberToPaymentParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createMemberToPayment, arg.EventID, arg.MemberID, arg.PaymentID)
}

const listMemberToPaymentsByEventID = `-- name: ListMemberToPaymentsByEventID :many
SELECT event_id, member_id, payment_id FROM member_to_payment WHERE event_id = ?
`

func (q *Queries) ListMemberToPaymentsByEventID(ctx context.Context, eventID string) ([]MemberToPayment, error) {
	rows, err := q.db.QueryContext(ctx, listMemberToPaymentsByEventID, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MemberToPayment
	for rows.Next() {
		var i MemberToPayment
		if err := rows.Scan(&i.EventID, &i.MemberID, &i.PaymentID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
