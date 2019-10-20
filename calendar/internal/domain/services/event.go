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
