package usecase

import (
	"bspliter/domain/model"
	"bspliter/infra"
	"context"
)

type EventController interface {
	Create(ctx context.Context, name string, memberNames []string) (*model.Event, error)
}

type EventControllerImpl struct {
	eventRepository infra.EventRepository
}

func NewEventController(eventRepository infra.EventRepository) EventController {
	return &EventControllerImpl{eventRepository: eventRepository}
}

func (c *EventControllerImpl) Create(ctx context.Context, name string, memberNames []string) (*model.Event, error) {
	event, err := model.NewEvent(name, memberNames)
	if err != nil {
		return nil, err
	}
	if err := c.eventRepository.Store(ctx, event); err != nil {
		return nil, err
	}
	return event, nil
}
