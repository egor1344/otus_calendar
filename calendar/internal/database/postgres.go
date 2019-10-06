package database

import (
	"context"
	"log"

	"github.com/egor1344/otus_calendar/calendar/internal/domain/models"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// PgEventStorage implements domain/interfaces/database
type PgEventStorage struct {
	db *sqlx.DB
}

// NewPgEventStorage init storage
func NewPgEventStorage(dsn string) (*PgEventStorage, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgEventStorage{db: db}, nil
}

// AddEvent save event
func (pges *PgEventStorage) AddEvent(ctx context.Context, event *models.Event) error {
	log.Println("AddEvent!")
	query := `
		INSERT INTO events(title, date_time, duration, owner, description)
		VALUES (:title, :date_time, :duration, :owner, :description)
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"title":       event.Title,
		"date_time":   event.Datetime,
		"duration":    event.Duration,
		"owner":       event.UserID,
		"description": event.Description,
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// GetEventByID get event
func (pges *PgEventStorage) GetEventByID(ctx context.Context, id string) (*models.Event, error) {
	log.Println("GetEventByID")
	query := `
		SELECT (title, date_time, duration, owner, description) FROM events
		WHERE (id=:id)
	`
	row, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id": id,
	})
	log.Println(row)
	if err != nil {
		log.Fatal(err)
	}
	return nil, nil
}

// UpdateEventByID update event
func (pges *PgEventStorage) UpdateEventByID(ctx context.Context, id int64, updateEvent *models.Event) (*models.Event, error) {
	log.Println("UpdateEventByID")
	return nil, nil
}

// DeleteEventByID delete event
func (pges *PgEventStorage) DeleteEventByID(ctx context.Context, id int64) error {
	log.Println("DeleteEventByID")
	return nil
}
