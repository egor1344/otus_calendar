package models

import (
	"time"
)

// Event - структура события
type Event struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Datetime    time.Time `json:"date_time" db:"date_time"`
	Duration    int64     `json:"duration" db:"duration"`
	Description string    `json:"description" db:"description"`
	UserID      int64     `json:"ower" db:"owner"`
}
