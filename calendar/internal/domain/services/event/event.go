package event

import (
	"context"
	"time"

	interfaces "github.com/egor1344/otus_calendar/calendar/internal/interfaces/database"
	logger "github.com/egor1344/otus_calendar/calendar/pkg/logger"
)

// Service сервис предоставляющий работу с событиями
type Service struct {
	Database interfaces.Database
}

// AddEvent добавление события
func (e *Service) AddEvent(ctx *context.Context, title string, datatime *time.Time, duration int64, description string, userID int64) {
	logger.GetLogger()
	logger.Infof(ctx, title, datatime, duration, description, userID)
}
