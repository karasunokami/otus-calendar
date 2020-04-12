//+build wireinject

package amqp

import (
	"github.com/google/wire"
	"github.com/karasunokami/otus-calendar/internal/logger"
)

func InitializeAmqp() (*MessageBus, error) {
	wire.Build(NewMessageBus, CreateConfigFromEnvironment, logger.NewLogger)

	return &MessageBus{}, nil
}
