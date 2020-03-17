package db

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/dal"
	"time"
)

type Client interface {
	Create(event dal.Event) (EventId, error)

	Delete(id EventId) error

	Update(id EventId, event dal.Event) error

	Get(id EventId) (dal.Event, error)
}

type clientImpl struct {
	events map[EventId]dal.Event
}

type EventId int

func NewClient() Client {
	return &clientImpl{events: make(map[EventId]dal.Event)}
}

func (c *clientImpl) Get(id EventId) (dal.Event, error) {
	if !c.eventExists(id) {
		return dal.Event{}, EventNotFoundError
	}

	return c.events[id], nil
}

func (c *clientImpl) Create(event dal.Event) (EventId, error) {
	if busy := !c.isTimeAvailable(event.StartDatetime); busy {
		return 0, TimeBusyError
	}

	eventId := EventId(len(c.events) + 1)
	c.events[eventId] = event

	return eventId, nil
}

func (c *clientImpl) Update(id EventId, event dal.Event) error {
	if !c.eventExists(id) {
		return EventNotFoundError
	}

	if busy := !c.isTimeAvailable(event.StartDatetime); busy {
		return TimeBusyError
	}

	c.events[id] = event
	return nil
}

func (c *clientImpl) Delete(id EventId) error {
	if !c.eventExists(id) {
		return EventNotFoundError
	}

	delete(c.events, id)

	return nil
}

func (c clientImpl) isTimeAvailable(time time.Time) bool {
	for _, evt := range c.events {
		if evt.StartDatetime == time {
			return false
		}
	}

	return true
}

func (c clientImpl) eventExists(id EventId) bool {
	_, ok := c.events[id]

	return ok
}
