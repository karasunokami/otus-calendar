// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package amqp

import (
	"github.com/karasunokami/otus-calendar/internal/logger"
)

// Injectors from wire.go:

func InitializeAmqp() (*MessageBus, error) {
	amqpConfig := CreateConfigFromEnvironment()
	logrusLogger := logger.NewLogger()
	messageBus, err := NewMessageBus(amqpConfig, logrusLogger)
	if err != nil {
		return nil, err
	}
	return messageBus, nil
}
