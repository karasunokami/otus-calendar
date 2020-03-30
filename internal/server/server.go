package server

import (
	"context"
	"github.com/karasunokami/go.otus.hw/calendar/internal/calendar"
	"github.com/karasunokami/go.otus.hw/calendar/internal/config"
	api "github.com/karasunokami/go.otus.hw/calendar/pkg/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Server interface {
	Run() error

	CreateEvent(_ context.Context, req *api.CreateEventRequest) (*api.CreateEventResponse, error)
	UpdateEvent(_ context.Context, req *api.UpdateEventRequest) (*api.UpdateEventResponse, error)
	GetEvent(_ context.Context, req *api.GetEventRequest) (*api.GetEventResponse, error)
	DeleteEvent(_ context.Context, req *api.DeleteEventRequest) (*api.DeleteEventResponse, error)
}

type serverImpl struct {
	calendar calendar.Calendar
	conf     config.HttpListen
	logger   logrus.FieldLogger

	server *grpc.Server
}

func NewHttpServer(calendar calendar.Calendar, logger logrus.FieldLogger, conf config.HttpListen) Server {
	return &serverImpl{calendar: calendar, logger: logger, conf: conf}
}

func (s *serverImpl) Run() error {
	conn, err := net.Listen("tcp", net.JoinHostPort(s.conf.Ip, s.conf.Port))
	if err != nil {
		s.logger.WithError(err).Fatal("Failed to start api listener")
		return err
	}

	s.server = grpc.NewServer(grpc.ConnectionTimeout(s.conf.ConnTimeout))
	api.RegisterCalendarServiceServer(s.server, s)

	if err := s.server.Serve(conn); err != nil {
		s.logger.WithError(err).Fatal("Failed to start api server")
	}

	return nil
}
