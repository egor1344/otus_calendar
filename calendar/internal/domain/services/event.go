package services

import (
	"context"
	"log"

	interfaces "github.com/egor1344/otus_calendar/calendar/internal/domain/interfaces"
	logger "github.com/egor1344/otus_calendar/calendar/pkg/logger"
	models "github.com/egor1344/otus_calendar/internal/domain/models"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// Service сервис предоставляющий работу с событиями
type Service struct {
	Database interfaces.Database
}

// AddEvent добавление события
func (e *Service) AddEvent(ctx context.Context, title string, date *timestamp.Timestamp, duration int64, description string, userID int64) (models.Event, error) {
	zapLog, err := logger.GetLogger()
	if err != nil {
		log.Fatal(err)
	}
	zapLog.Info(ctx, title, date, duration, description, userID)
	return nil, nil
}
