package database

import (
	"context"
	"github.com/egor1344/otus_calendar/calendar/internal/domain/models"
	"github.com/egor1344/otus_calendar/calendar/proto/event"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"log"
	"time"
)

// PgEventStorage implements domain/interfaces/database
type PgEventStorage struct {
	db  *sqlx.DB
	Log *zap.SugaredLogger
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
func (pges *PgEventStorage) AddEvent(ctx context.Context, event *event.Event) (string, error) {
	pges.Log.Info("AddEvent!")
	dateTime, err := ptypes.Timestamp(event.Datetime)
	if err != nil {
		log.Fatal(err)
	}
	query := `
		INSERT INTO events(title, date_time, duration, owner, description, before_time_pull, send)
		VALUES (:title, :date_time, :duration, :owner, :description, :before_time_pull, :send) RETURNING id
	`
	result, err := pges.db.NamedQueryContext(ctx, query, map[string]interface{}{
		"title":            event.Title,
		"date_time":        dateTime,
		"duration":         event.Duration,
		"owner":            event.UserId,
		"description":      event.Description,
		"before_time_pull": event.BeforeTimePull,
		"send":             false,
	})
	if err != nil {
		pges.Log.Fatal(err)
	}
	defer result.Close()

	var returnUUID string
	result.Next()
	result.Scan(&returnUUID)
	pges.Log.Info(result, returnUUID)
	return returnUUID, nil
}

// GetEventByID get event
func (pges *PgEventStorage) GetEventByID(ctx context.Context, id string) (*event.Event, error) {
	pges.Log.Info("GetEventByID")
	ev := event.Event{}
	row, err := pges.db.Query(`SELECT id, title, date_time, duration, owner, description FROM events WHERE id::uuid in ($1)`, id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if row.Next() {
		var datetime time.Time
		err = row.Scan(&ev.Uuid, &ev.Title, &datetime, &ev.Duration, &ev.UserId, &ev.Description)
		if err != nil {
			log.Fatal(err)
		}
		ev.Datetime, err = ptypes.TimestampProto(datetime)
		if err != nil {
			log.Fatal(err)
		}
	}
	pges.Log.Info(ev)
	return &ev, nil
}

// UpdateEventByID update event
func (pges *PgEventStorage) UpdateEventByID(ctx context.Context, id int64, updateEvent *models.Event) (*models.Event, error) {
	pges.Log.Info("UpdateEventByID")
	return nil, nil
}

// DeleteEventByID delete event
func (pges *PgEventStorage) DeleteEventByID(ctx context.Context, id int64) error {
	pges.Log.Info("DeleteEventByID")
	return nil
}
