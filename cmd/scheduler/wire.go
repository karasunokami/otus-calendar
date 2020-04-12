//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/karasunokami/otus-calendar/internal/domain"
	"github.com/karasunokami/otus-calendar/internal/logger"
	"github.com/karasunokami/otus-calendar/internal/messaging/amqp"
	"github.com/karasunokami/otus-calendar/internal/service"
	"github.com/karasunokami/otus-calendar/internal/storage/sql"
)

func InitializeScheduler() (*scheduler, error) {
	wire.Build(
		newScheduler,
		sql.InitializeStorage,
		amqp.InitializeAmqp,
		wire.Bind(new(domain.StorageInterface), new(*sql.Storage)),
		wire.Bind(new(domain.MessageBusInterface), new(*amqp.MessageBus)),
		service.NewCalendarService,
		logger.NewLogger,
	)

	return &scheduler{}, nil
}
