package calendar

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/dal"
	"github.com/karasunokami/go.otus.hw/calendar/internal/db"
	"github.com/sirupsen/logrus"
	"time"
)

type Calendar interface {
	GetEvent(id string) (dal.Event, error)

	CreateEvent(title string, startTime time.Time, endTime time.Time) (*dal.Event, error)

	UpdateEvent(event dal.Event) error

	DeleteEvent(id string) error
}

type calendarImpl struct {
	client db.Client
	logger logrus.FieldLogger
}

func NewCalendar(client db.Client, logger logrus.FieldLogger) *calendarImpl {
	return &calendarImpl{client: client, logger: logger}
}

func (s *calendarImpl) CreateEvent(title string, startTime time.Time, endTime time.Time) (*dal.Event, error) {
	evt := dal.Event{
		Title:     title,
		StartTime: startTime,
		EndTime:   endTime,
	}

	event, err := s.client.Create(&evt)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *calendarImpl) GetEvent(id string) (dal.Event, error) {
	evt, err := s.client.Get(id)
	if err != nil {
		return evt, err
	}

	return evt, nil
}

func (s *calendarImpl) UpdateEvent(event dal.Event) error {
	return s.client.Update(event)
}

func (s *calendarImpl) DeleteEvent(id string) error {
	return s.client.Delete(id)
}
