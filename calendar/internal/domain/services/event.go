package services

import (
	"context"
	"github.com/egor1344/otus_calendar/calendar/internal/domain/interfaces"
	"github.com/egor1344/otus_calendar/calendar/proto/event"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.uber.org/zap"
)

// Service сервис предоставляющий работу с событиями
type Service struct {
	Database interfaces.Database
	Log      *zap.SugaredLogger
}

// AddEvent добавление события
func (e *Service) AddEvent(ctx context.Context, title string, date *timestamp.Timestamp, duration int64, description string, userID int64, beforeTimePull int64) (*event.Event, error) {

	ev := event.Event{Title: title, Datetime: date, Duration: duration, Description: description, UserId: userID, BeforeTimePull: beforeTimePull}
	id, err := e.Database.AddEvent(ctx, &ev)
	if err != nil {
		e.Log.Fatal(err)
	}
	e.Log.Info(id)
	ev.Uuid = id
	e.Log.Info(ev)
	return &ev, nil
}

// GetEvent получение события
func (e *Service) GetEvent(ctx context.Context, uuid string) (*event.Event, error) {
	ev, err := e.Database.GetEventByID(ctx, uuid)
	if err != nil {
		e.Log.Fatal(err)
	}
	e.Log.Info(ev)
	return ev, nil
}

// UpdateEvent обновление событий
func (e *Service) UpdateEvent(ctx context.Context, event *event.Event) (*event.Event, error) {
	ev, err := e.Database.UpdateEventByID(ctx, event)
	if err != nil {
		e.Log.Fatal(err)
	}
	e.Log.Info(ev)
	return ev, nil
}

// DeleteEvent удалений событий
func (e *Service) DeleteEvent(ctx context.Context, id string) error {
	err := e.Database.DeleteEventByID(ctx, id)
	if err != nil {
		e.Log.Fatal(err)
	}
	return nil
}

// GetEventList получение списка событий по типу и userID
func (e *Service) GetEventList(ctx context.Context, types string, userId string) (events []*event.Event, err error) {
	events, err = e.Database.GetEventList(ctx, types, userId)
	if err != nil {
		e.Log.Fatal(err)
	}
	return events, nil
}
