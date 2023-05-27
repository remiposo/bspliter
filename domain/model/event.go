package model

import (
	"fmt"
	"unicode/utf8"

	"github.com/google/uuid"
)

type Event struct {
	ID       string
	Name     string
	Members  []*Member
	Payments []*Payment
}

type Member struct {
	ID   string
	Name string
}

type Payment struct {
	ID     string
	Name   string
	Amount int
	Payer  string
	Payees []string
}

func NewEvent(name string, memberNames []string) (*Event, error) {
	// validate name
	if name == "" || utf8.RuneCountInString(name) > 255 {
		return nil, fmt.Errorf("name is empty or too long")
	}

	// validate memberNames
	mMap := make(map[string]struct{})
	for _, mName := range memberNames {
		if mName == "" || utf8.RuneCountInString(mName) > 255 {
			return nil, fmt.Errorf("member name is empty or too long")
		}
		if _, ok := mMap[mName]; ok {
			return nil, fmt.Errorf("member name %s is duplicated", mName)
		}
		mMap[mName] = struct{}{}
	}

	event := &Event{
		ID:       uuid.New().String(),
		Name:     name,
		Members:  make([]*Member, 0, len(memberNames)),
		Payments: make([]*Payment, 0),
	}
	for _, mName := range memberNames {
		event.Members = append(event.Members, &Member{
			ID:   uuid.New().String(),
			Name: mName,
		})
	}
	return event, nil
}

func ReconstructEvent(id, name string, members []*Member, payments []*Payment) *Event {
	return &Event{
		ID:       id,
		Name:     name,
		Members:  members,
		Payments: payments,
	}
}

func ReconstructMember(id, name string) *Member {
	return &Member{
		ID:   id,
		Name: name,
	}
}

func ReconstructPayment(id, name string, amount int, payer string, payees []string) *Payment {
	return &Payment{
		ID:     id,
		Name:   name,
		Amount: amount,
		Payer:  payer,
		Payees: payees,
	}
}

func (e *Event) AddPayment(name string, amount int, payer string, payees []string) error {
	// validate name
	if name == "" || utf8.RuneCountInString(name) > 255 {
		return fmt.Errorf("name is empty or too long")
	}

	// validate amount
	if amount < 0 {
		return fmt.Errorf("amount is not positive")
	}

	// validate payer
	pFound := false
	for _, member := range e.Members {
		if member.ID == payer {
			pFound = true
			break
		}
	}
	if !pFound {
		return fmt.Errorf("payer %s is not found in members", payer)
	}

	// validate payees
	if len(payees) == 0 {
		return fmt.Errorf("payees is empty")
	}
	pMap := make(map[string]struct{})
	for _, payee := range payees {
		if _, ok := pMap[payee]; ok {
			return fmt.Errorf("payee name %s is duplicated", payee)
		}
		pMap[payee] = struct{}{}
		pFound := false
		for _, member := range e.Members {
			if member.ID == payee {
				pFound = true
				break
			}
		}
		if !pFound {
			return fmt.Errorf("payee %s is not found in members", payee)
		}
	}

	e.Payments = append(e.Payments, &Payment{
		ID:     uuid.New().String(),
		Name:   name,
		Amount: amount,
		Payer:  payer,
		Payees: payees,
	})
	return nil
}
