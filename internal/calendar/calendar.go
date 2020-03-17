package calendar

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/dal"
	"github.com/karasunokami/go.otus.hw/calendar/internal/db"
	"github.com/sirupsen/logrus"
	"time"
)

type Calendar interface {
	GetEvent(id db.EventId) (dal.Event, error)

	CreateEvent(startTime time.Time, stopTime time.Time) (db.EventId, error)

	UpdateEvent(id db.EventId, event dal.Event) error

	DeleteEvent(id db.EventId) error
}

type calendarImpl struct {
	client db.Client
	logger logrus.FieldLogger
}

func NewCalendar(client db.Client, logger logrus.FieldLogger) *calendarImpl {
	return &calendarImpl{client: client, logger: logger}
}

func (s *calendarImpl) CreateEvent(startDatetime, endDatetime time.Time) (db.EventId, error) {
	logFields := logrus.Fields{
		"module": "calendar",
		"util":   "calendar",
		"cmp":    "CreateEvent",
	}

	evt := dal.Event{
		Title:         "",
		StartDatetime: startDatetime,
		EndDatetime:   endDatetime,
	}

	id, err := s.client.Create(evt)
	if err != nil {
		s.logger.WithFields(logFields).WithError(err).Error("Failed to create event")
	}

	return id, nil
}

func (s *calendarImpl) GetEvent(id db.EventId) (dal.Event, error) {
	logFields := logrus.Fields{
		"module": "calendar",
		"util":   "calendar",
		"cmp":    "GetEvent",
	}

	evt, err := s.client.Get(id)
	if err != nil {
		s.logger.WithFields(logFields).WithError(err).Error("Failed to get event")
	}

	return evt, nil
}

func (s *calendarImpl) UpdateEvent(id db.EventId, event dal.Event) error {
	return s.client.Update(id, event)
}

func (s *calendarImpl) DeleteEvent(id db.EventId) error {
	return s.client.Delete(id)
}
