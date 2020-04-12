package handler

import (
	"errors"
	"net/http"
	"time"
)

type eventCreateRequest struct {
	Title    string
	DateFrom time.Time
	DateTo   time.Time
}

type eventCreateResponse struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}

// EventCreateHandler Add new event to storage
// POST /create_event
// Fields:
// 	 title:     string
//   date_from: datetime, format 2019-01-02T12:00:00Z
//   date_to:   datetime, format 2019-01-02T14:00:00Z
func (h *Handler) EventCreateHandler(req *http.Request) APIResponse {
	data := &eventCreateRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	event, err := h.Calendar.CreateEvent(data.Title, data.DateFrom, data.DateTo)
	if err != nil {
		h.Error(http.StatusInternalServerError, err)
	}

	return h.JSON(http.StatusCreated, &eventCreateResponse{
		ID:       event.ID.String(),
		Title:    event.Title,
		DateFrom: event.DateFrom.Format(time.RFC3339),
		DateTo:   event.DateTo.Format(time.RFC3339),
	})
}

func (data *eventCreateRequest) parse(req *http.Request) error {
	var err error
	if err := req.ParseForm(); err != nil {
		return err
	}

	data.Title = req.FormValue("title")
	if data.Title == "" {
		return errors.New("event title is required")
	}

	if data.DateFrom, err = time.Parse(time.RFC3339, req.FormValue("date_from")); err != nil {
		return err
	}

	if data.DateTo, err = time.Parse(time.RFC3339, req.FormValue("date_to")); err != nil {
		return err
	}

	if data.DateFrom.After(data.DateTo) {
		return errors.New("date From can not be after Date To")
	}

	return nil
}
