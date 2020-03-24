package db

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/dal"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Client interface {
	Create(event *dal.Event) (*dal.Event, error)

	Delete(id string) error

	Update(event dal.Event) error

	Get(id string) (dal.Event, error)
}

type clientImpl struct {
	events map[string]dal.Event
}

func NewClient() Client {
	return &clientImpl{events: make(map[string]dal.Event)}
}

func (c *clientImpl) Get(id string) (dal.Event, error) {
	if !c.eventExists(id) {
		return dal.Event{}, EventNotFoundError
	}

	return c.events[id], nil
}

func (c *clientImpl) Create(event *dal.Event) (*dal.Event, error) {
	if busy := !c.isTimeAvailable(event.StartTime); busy {
		return (*dal.Event)(nil), TimeBusyError
	}

	event.ID = uuid.NewV4()
	c.events[event.ID.String()] = *event

	return event, nil
}

func (c *clientImpl) Update(event dal.Event) error {
	fetched, err := c.Get(event.ID.String())
	if err != nil {
		return err
	}

	if fetched.StartTime != event.StartTime || fetched.EndTime != fetched.EndTime {
		if busy := !c.isTimeAvailable(event.StartTime); busy {
			return TimeBusyError
		}
	}

	c.events[event.ID.String()] = event

	return nil
}

func (c *clientImpl) Delete(id string) error {
	if !c.eventExists(id) {
		return EventNotFoundError
	}

	delete(c.events, id)

	return nil
}

func (c clientImpl) isTimeAvailable(time time.Time) bool {
	for _, evt := range c.events {
		if evt.StartTime == time {
			return false
		}
	}

	return true
}

func (c clientImpl) eventExists(id string) bool {
	_, ok := c.events[id]

	return ok
}
