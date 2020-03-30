package server

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/app"
	"github.com/karasunokami/go.otus.hw/calendar/internal/calendar"
	"github.com/karasunokami/go.otus.hw/calendar/internal/config"
	"github.com/sirupsen/logrus"
)

type Provider struct{}

func (p *Provider) Register(c *app.Container) {
	c.Set("server", func(c *app.Container) interface{} {
		calendarSvc := c.Get("calendar").(calendar.Calendar)
		conf := c.Get("config").(*config.AppConfig)
		logger := c.Get("logger").(logrus.FieldLogger)

		return NewHttpServer(calendarSvc, logger, conf.HttpListen)
	})
}
