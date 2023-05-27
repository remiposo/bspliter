package presentation

import (
	"bspliter/usecase"

	"github.com/labstack/echo/v4"
)

type EventHandler interface {
	Create(ctx echo.Context) error
}

type EventHandlerImpl struct {
	eventController usecase.EventController
}

func NewEventHandler(eventController usecase.EventController) EventHandler {
	return &EventHandlerImpl{eventController: eventController}
}

func (h *EventHandlerImpl) Create(ctx echo.Context) error {
	type Req struct {
		Name        string   `json:"name"`
		MemberNames []string `json:"member_names"`
	}
	req := new(Req)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	event, err := h.eventController.Create(ctx.Request().Context(), req.Name, req.MemberNames)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	type RespMember struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	type Resp struct {
		ID      string        `json:"id"`
		Name    string        `json:"name"`
		Members []*RespMember `json:"members"`
	}
	resp := &Resp{
		ID:      event.ID,
		Name:    event.Name,
		Members: make([]*RespMember, len(event.Members)),
	}
	for i, member := range event.Members {
		resp.Members[i] = &RespMember{
			ID:   member.ID,
			Name: member.Name,
		}
	}
	return ctx.JSON(200, resp)
}
