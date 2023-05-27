package infra

import (
	"bspliter/domain/model"
	"bspliter/infra/db"
	"context"
	"database/sql"
)

type EventRepository interface {
	Get(ctx context.Context, id string) (*model.Event, error)
	Store(ctx context.Context, event *model.Event) error
}

type EventRepositoryImpl struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{db: db}
}

func (r *EventRepositoryImpl) Get(ctx context.Context, id string) (*model.Event, error) {
	q := db.New(r.db)
	eventDTOs, err := q.GetEventByID(ctx, id)
	if err != nil {
		return nil, err
	}
	memberDTOs, err := q.ListMembersByEventID(ctx, id)
	if err != nil {
		return nil, err
	}
	paymentDTOs, err := q.ListPaymentsByEventID(ctx, id)
	if err != nil {
		return nil, err
	}
	memberToPaymentDTOs, err := q.ListMemberToPaymentsByEventID(ctx, id)
	if err != nil {
		return nil, err
	}

	// reconstruct event from DTOs
	payeeMap := make(map[string][]string)
	for _, mtp := range memberToPaymentDTOs {
		payeeMap[mtp.PaymentID] = append(payeeMap[mtp.PaymentID], mtp.MemberID)
	}
	members := make([]*model.Member, len(memberDTOs))
	for idx, memberDTO := range memberDTOs {
		members[idx] = model.ReconstructMember(memberDTO.ID, memberDTO.Name)
	}
	payments := make([]*model.Payment, len(paymentDTOs))
	for idx, paymentDTO := range paymentDTOs {
		payments[idx] = model.ReconstructPayment(paymentDTO.ID, paymentDTO.Name, int(paymentDTO.Amount), paymentDTO.MemberID, payeeMap[paymentDTO.ID])
	}
	return model.ReconstructEvent(eventDTOs.ID, eventDTOs.Name, members, payments), nil
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
			if _, err := q.CreateMemberToPayment(ctx, db.CreateMemberToPaymentParams{EventID: event.ID, PaymentID: payment.ID, MemberID: payee}); err != nil {
				return err
			}
		}
	}
	return tx.Commit()
}
