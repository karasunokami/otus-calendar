package dal

import (
	"github.com/satori/go.uuid"
	"time"
)

type Event struct {
	ID        uuid.UUID
	Title     string
	StartTime time.Time
	EndTime   time.Time
}
