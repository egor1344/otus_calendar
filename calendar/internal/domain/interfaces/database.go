package interfaces

import (
	"context"
	"github.com/egor1344/otus_calendar/calendar/proto/event"

	"github.com/egor1344/otus_calendar/calendar/internal/domain/models"
)

// Database - Интерфейс реализующий работу с бд
type Database interface {
	AddEvent(ctx context.Context, event *event.Event) (string, error)
	GetEventByID(ctx context.Context, id string) (*event.Event, error)
	UpdateEventByID(ctx context.Context, id int64, updateEvent *models.Event) (*models.Event, error)
	DeleteEventByID(ctx context.Context, id int64) error
}
