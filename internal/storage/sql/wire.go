//+build wireinject

package sql

import (
	"github.com/google/wire"
	"github.com/karasunokami/otus-calendar/internal/logger"
)

func InitializeStorage() (*Storage, error) {
	wire.Build(NewStorage, CreateConfigFromEnvironment, logger.NewLogger)

	return &Storage{}, nil
}
