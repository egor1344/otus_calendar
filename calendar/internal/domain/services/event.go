package services

import (
	"context"
	"log"

	interfaces "github.com/egor1344/otus_calendar/calendar/internal/domain/interfaces"
	models "github.com/egor1344/otus_calendar/calendar/internal/domain/models"
	logger "github.com/egor1344/otus_calendar/calendar/pkg/logger"
	ptypes "github.com/golang/protobuf/ptypes"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// Service сервис предоставляющий работу с событиями
type Service struct {
	Database interfaces.Database
}

// AddEvent добавление события
func (e *Service) AddEvent(ctx context.Context, title string, date *timestamp.Timestamp, duration int64, description string, userID int64) (*models.Event, error) {
	zapLog, err := logger.GetLogger()
	if err != nil {
		log.Fatal(err)
	}
	// todo: validation
	dateTime, err := ptypes.Timestamp(date)
	if err != nil {
		log.Fatal(err)
	}
	event := models.Event{Title: title, Datetime: dateTime, Duration: duration, Description: description, UserID: userID}
	zapLog.Info(event)
	e.Database.AddEvent(ctx, &event)
	return &event, nil
}
