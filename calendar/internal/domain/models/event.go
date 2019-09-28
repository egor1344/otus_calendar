package models

import (
	"time"
)

// Event - структура события
type Event struct {
	ID          int64
	Title       string
	Datetime    time.Time
	Duration    int64
	Description string
	UserID      int64
}
