package presentation

import (
	"bspliter/domain/model"
	"bspliter/usecase"

	"github.com/labstack/echo/v4"
)

type EventHandler interface {
	Create(c echo.Context) error
	AddPayment(c echo.Context) error
}

type EventHandlerImpl struct {
	eventController usecase.EventController
}

func NewEventHandler(eventController usecase.EventController) EventHandler {
	return &EventHandlerImpl{eventController: eventController}
}

func (h *EventHandlerImpl) Create(c echo.Context) error {
	type Req struct {
		Name        string   `json:"name"`
		MemberNames []string `json:"member_names"`
	}
	req := new(Req)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	event, err := h.eventController.Create(c.Request().Context(), req.Name, req.MemberNames)
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
	return h.respEvent(c, event)
}

func (h *EventHandlerImpl) AddPayment(c echo.Context) error {
	type Req struct {
		EventID  string   `param:"id"`
		Name     string   `json:"name"`
		Amount   int      `json:"amount"`
		PayerID  string   `json:"payer_id"`
		PayeeIDs []string `json:"payee_ids"`
	}
	req := new(Req)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	event, err := h.eventController.AddPayment(c.Request().Context(), req.EventID, req.Name, req.Amount, req.PayerID, req.PayeeIDs)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	type Resp struct {
		ID       string   `json:"id"`
		Name     string   `json:"name"`
		Amount   int      `json:"amount"`
		PayerID  string   `json:"payer_id"`
		PayeeIDs []string `json:"payee_ids"`
	}
	return h.respEvent(c, event)
}

func (h *EventHandlerImpl) respEvent(c echo.Context, event *model.Event) error {
	type RespMember struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	type RespPayment struct {
		ID       string   `json:"id"`
		Name     string   `json:"name"`
		Amount   int      `json:"amount"`
		PayerID  string   `json:"payer_id"`
		PayeeIDs []string `json:"payee_ids"`
	}
	type Resp struct {
		ID       string         `json:"id"`
		Name     string         `json:"name"`
		Members  []*RespMember  `json:"members"`
		Payments []*RespPayment `json:"payments"`
	}
	resp := &Resp{
		ID:       event.ID,
		Name:     event.Name,
		Members:  make([]*RespMember, len(event.Members)),
		Payments: make([]*RespPayment, len(event.Payments)),
	}
	for i, member := range event.Members {
		resp.Members[i] = &RespMember{
			ID:   member.ID,
			Name: member.Name,
		}
	}
	for i, payment := range event.Payments {
		resp.Payments[i] = &RespPayment{
			ID:       payment.ID,
			Name:     payment.Name,
			Amount:   payment.Amount,
			PayerID:  payment.Payer,
			PayeeIDs: payment.Payees,
		}
	}
	return c.JSON(200, resp)
}
