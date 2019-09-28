package interfaces

import (
	"context"
	"github.com/egor1344/otus_calendar/calendar/internal/domain/models"
)

// Database - Интерфейс реализующий работу с бд
type Database interface {
	SaveEvent(ctx context.Context, event *models.Event) error
	GetEventByID(ctx context.Context, id int64) (*models.Event, error)
	UpdateEventByID(ctx context.Context, id int64, updateEvent *models.Event) (*models.Event, error)
	DeleteEventByID(ctx context.Context, id int64) error
}
