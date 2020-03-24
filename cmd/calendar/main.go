package main

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/api/server"
	"github.com/karasunokami/go.otus.hw/calendar/internal/app"
	"github.com/karasunokami/go.otus.hw/calendar/internal/calendar"
	"github.com/karasunokami/go.otus.hw/calendar/internal/config"
	"github.com/karasunokami/go.otus.hw/calendar/internal/db"
	"github.com/karasunokami/go.otus.hw/calendar/internal/logger"
)

func main() {
	cont := app.NewContainer()

	cont.Register(new(config.Provider))
	cont.Register(new(logger.Provider))
	cont.Register(new(db.Provider))
	cont.Register(new(calendar.Provider))
	cont.Register(new(server.Provider))

	_ = cont.Get("server").(server.Server).Run()
}
