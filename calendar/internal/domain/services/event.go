package services

import (
	"context"
	"log"

	"github.com/egor1344/otus_calendar/calendar/internal/domain/interfaces"
	"github.com/egor1344/otus_calendar/calendar/internal/domain/models"
	"github.com/egor1344/otus_calendar/calendar/pkg/logger"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// Service сервис предоставляющий работу с событиями
type Service struct {
	Database interfaces.Database
}

// AddEvent добавление события
func (e *Service) AddEvent(ctx context.Context, title string, date *timestamp.Timestamp, duration int64, description string, userID int64, beforeTimePull int64) (*models.Event, error) {
	zapLog, err := logger.GetLogger()
	if err != nil {
		log.Fatal(err)
	}
	// todo: validation
	dateTime, err := ptypes.Timestamp(date)
	if err != nil {
		log.Fatal(err)
	}
	event := models.Event{Title: title, DateTime: dateTime, Duration: duration, Description: description, UserID: userID, BeforeTimePull: beforeTimePull}
	id, err := e.Database.AddEvent(ctx, &event)
	if err != nil {
		log.Fatal(err)
	}
	zapLog.Info(id)
	event.UUID = id
	zapLog.Info(event)
	return &event, nil
}
