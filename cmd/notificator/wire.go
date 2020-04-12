//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/karasunokami/otus-calendar/internal/domain"
	"github.com/karasunokami/otus-calendar/internal/logger"
	"github.com/karasunokami/otus-calendar/internal/messaging/amqp"
)

func InitializeNotificator() (*notificator, error) {
	wire.Build(
		newNotificator,
		amqp.InitializeAmqp,
		wire.Bind(new(domain.MessageBusInterface), new(*amqp.MessageBus)),
		logger.NewLogger,
	)

	return &notificator{}, nil
}
