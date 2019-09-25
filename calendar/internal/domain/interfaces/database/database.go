package database

import (
	"context"
	"github.com/egor1344/otus_calendar/calendar/internal/model/event"
)

// Database - Интерфейс реализующий работу с бд
type Database interface {
	SaveEvent(ctx context.Context, event *event.Event) error
	GetEventByID(ctx context.Context, id int64) (*event.Event, error)
	UpdateEventByID(ctx context.Context, id int64, updateEvent *event.Event) (*event.Event, error)
	DeleteEventByID(ctx context.Context, id int64) error
}
