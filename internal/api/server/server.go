package server

import (
	"context"
	api "github.com/karasunokami/go.otus.hw/calendar/internal/api/proto"
	"github.com/karasunokami/go.otus.hw/calendar/internal/calendar"
	"github.com/karasunokami/go.otus.hw/calendar/internal/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Server interface {
	Run() error

	CreateEvent(_ context.Context, req *api.CreateEventRequest) (*api.CreateEventResponse, error)
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
