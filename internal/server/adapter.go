package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/karasunokami/go.otus.hw/calendar/internal/dal"
	"github.com/karasunokami/go.otus.hw/calendar/internal/db"
	api "github.com/karasunokami/go.otus.hw/calendar/pkg/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *serverImpl) UpdateEvent(_ context.Context, req *api.UpdateEventRequest) (*api.UpdateEventResponse, error) {
	logFields := logrus.Fields{
		"util":     "server",
		"cmp":      "update_event",
		"event_id": req.GetId(),
	}

	event, err := s.calendar.GetEvent(req.GetId())

	if err != nil {
		return nil, db.EventNotFoundError
	}

	event.Title = req.GetTitle()
	event.StartTime = time.Unix(req.StartTime.GetSeconds(), int64(req.StartTime.GetNanos()))
	event.EndTime = time.Unix(req.EndTime.GetSeconds(), int64(req.EndTime.GetNanos()))

	err = s.calendar.UpdateEvent(event)

	if err != nil {
		s.logger.WithFields(logFields).WithError(err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	s.logger.WithFields(logFields).Infof("event #%s updated", event.ID.String())

	return &api.UpdateEventResponse{Event: s.convertToProtoEvent(&event)}, nil
}

func (s *serverImpl) CreateEvent(_ context.Context, req *api.CreateEventRequest) (*api.CreateEventResponse, error) {
	logFields := logrus.Fields{
		"util": "server",
		"cmp":  "create_event",
	}

	event, err := s.calendar.CreateEvent(
		req.GetTitle(),
		time.Unix(req.StartTime.GetSeconds(), int64(req.StartTime.GetNanos())),
		time.Unix(req.EndTime.GetSeconds(), int64(req.EndTime.GetNanos())),
	)
	if err != nil {
		s.logger.WithFields(logFields).WithError(err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	s.logger.WithFields(logFields).Infof("event #%s created", event.ID.String())

	return &api.CreateEventResponse{Event: s.convertToProtoEvent(event)}, nil
}

func (s *serverImpl) GetEvent(_ context.Context, req *api.GetEventRequest) (*api.GetEventResponse, error) {
	logFields := logrus.Fields{
		"util": "server",
		"cmp":  "get_event",
		"id":   req.GetId(),
	}

	event, err := s.calendar.GetEvent(req.GetId())

	if err != nil {
		s.logger.WithFields(logFields).WithError(err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	s.logger.WithFields(logFields).Infof("event #%s fetched", event.ID.String())

	return &api.GetEventResponse{Event: s.convertToProtoEvent(&event)}, nil
}

func (s *serverImpl) DeleteEvent(_ context.Context, req *api.DeleteEventRequest) (*api.DeleteEventResponse, error) {
	logFields := logrus.Fields{
		"util": "server",
		"cmp":  "delete_event",
		"id":   req.GetId(),
	}

	err := s.calendar.DeleteEvent(req.GetId())

	if err != nil {
		s.logger.WithFields(logFields).WithError(err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	s.logger.WithFields(logFields).Infof("event #%s deleted", req.GetId())

	return &api.DeleteEventResponse{}, nil
}

func (s *serverImpl) convertToProtoEvent(event *dal.Event) *api.Event {
	return &api.Event{
		Id:        event.ID.String(),
		Title:     event.Title,
		StartTime: &timestamp.Timestamp{Seconds: event.StartTime.Unix(), Nanos: int32(event.StartTime.UnixNano())},
		EndTime:   &timestamp.Timestamp{Seconds: event.StartTime.Unix(), Nanos: int32(event.StartTime.UnixNano())},
	}
}
