package interfaces

import (
	"context"
	"github.com/egor1344/otus_calendar/calendar/proto/event"
)

// Database - Интерфейс реализующий работу с бд
type Database interface {
	AddEvent(ctx context.Context, event *event.Event) (string, error)
	GetEventByID(ctx context.Context, id string) (*event.Event, error)
	UpdateEventByID(ctx context.Context, updateEvent *event.Event) (*event.Event, error)
	DeleteEventByID(ctx context.Context, id string) error
	GetEventList(ctx context.Context, types string, userId string) ([]*event.Event, error)
}
