package models

import (
	"time"
)

// Event - структура события
type Event struct {
	UUID           string    `json:"id" db:"id"`
	Title          string    `json:"title" db:"title"`
	DateTime       time.Time `json:"date_time" db:"date_time"`
	Duration       int64     `json:"duration" db:"duration"`
	Description    string    `json:"description" db:"description"`
	UserID         int64     `json:"ower" db:"owner"`
	BeforeTimePull int64     `json:"before_time_pull" db:"before_time_pull"`
}
