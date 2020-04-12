//+build wireinject

package grpc

import (
	"github.com/google/wire"
	"github.com/karasunokami/otus-calendar/internal/domain"
	"github.com/karasunokami/otus-calendar/internal/logger"
	"github.com/karasunokami/otus-calendar/internal/service"
	"github.com/karasunokami/otus-calendar/internal/storage/sql"
)

func InitializeServer() (*server, error) {
	wire.Build(
		newServer,
		sql.InitializeStorage,
		service.NewCalendarService,
		wire.Bind(new(domain.StorageInterface), new(*sql.Storage)),
		logger.NewLogger,
	)

	return nil, nil
}
