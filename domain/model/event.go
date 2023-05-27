package model

import (
	"fmt"

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
	mMap := make(map[string]struct{})
	for _, mName := range memberNames {
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
