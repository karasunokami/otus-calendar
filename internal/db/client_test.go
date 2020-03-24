package db

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/dal"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ClientTestSuite struct {
	suite.Suite

	events []dal.Event
	client Client
}

func (s *ClientTestSuite) SetupTest() {

	timeFrom := time.Now()
	s.events = []dal.Event{
		{
			Title:     "1",
			StartTime: timeFrom,
			EndTime:   timeFrom.Add(time.Minute * 30),
		},
		{

			Title:     "2",
			StartTime: timeFrom.Add(time.Minute * 30 * 2),
			EndTime:   timeFrom.Add(time.Minute*30*2 + time.Minute*30),
		},
	}

	s.client = NewClient()
}

func (s *ClientTestSuite) TestClientCreate() {
	_, err := s.client.Create(&s.events[0])

	s.NoError(err)
}

func (s *ClientTestSuite) TestClientDelete() {
	evt, _ := s.client.Create(&s.events[0])
	err := s.client.Delete(evt.ID.String())

	s.NoError(err)
}

func (s *ClientTestSuite) TestClientUpdate() {
	evt, err := s.client.Create(&s.events[0])

	s.NoError(err)

	evt.Title = "new"
	err = s.client.Update(*evt)
	s.NoError(err)

	fetched, err := s.client.Get(evt.ID.String())

	s.NoError(err)

	s.Equal("new", fetched.Title)
}

func (s *ClientTestSuite) TestClientGet() {
	evt, _ := s.client.Create(&s.events[0])

	newEvt, err := s.client.Get(evt.ID.String())
	s.NoError(err)

	s.Equal(s.events[0].Title, newEvt.Title)
}

func (s *ClientTestSuite) TestCreateWithSameDate() {
	_, err := s.client.Create(&s.events[0])
	s.NoError(err)

	_, err = s.client.Create(&s.events[0])
	s.EqualError(err, "time is busy")
}

func (s *ClientTestSuite) TestGetNotFound() {
	_, err := s.client.Get("")
	s.EqualError(err, "dal not found")
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
