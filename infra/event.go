package infra

import (
	"bspliter/domain/model"
	"bspliter/infra/db"
	"context"
	"database/sql"
)

type EventRepository interface {
	Store(ctx context.Context, event *model.Event) error
}

type EventRepositoryImpl struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{db: db}
}

func (r *EventRepositoryImpl) Store(ctx context.Context, event *model.Event) error {
	// start transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	q := db.New(tx)
	defer tx.Rollback()

	// clean up all event related data
	if _, err := q.DeleteEvent(ctx, event.ID); err != nil {
		return err
	}

	// create all event related data
	if _, err := q.CreateEvent(ctx, db.CreateEventParams{ID: event.ID, Name: event.Name}); err != nil {
		return err
	}
	for _, member := range event.Members {
		if _, err := q.CreateMember(ctx, db.CreateMemberParams{ID: member.ID, Name: member.Name, EventID: event.ID}); err != nil {
			return err
		}
	}
	for _, payment := range event.Payments {
		if _, err := q.CreatePayment(ctx, db.CreatePaymentParams{ID: payment.ID, Name: payment.Name, Amount: int64(payment.Amount), EventID: event.ID, MemberID: payment.Payer}); err != nil {
			return err
		}
		for _, payee := range payment.Payees {
			if _, err := q.CreateMemberToPayment(ctx, db.CreateMemberToPaymentParams{PaymentID: payment.ID, MemberID: payee}); err != nil {
				return err
			}
		}
	}
	return nil
}
