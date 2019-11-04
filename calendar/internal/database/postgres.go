package database

import (
	"context"
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
	var returnUUID string
	result.Next()
	err = result.Scan(&returnUUID)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	log.Println(result, returnUUID)
	return returnUUID, nil
}

// GetEventByID get event
func (pges *PgEventStorage) GetEventByID(ctx context.Context, id string) (*event.Event, error) {
	pges.Log.Info("GetEventByID ", id)
	ev := event.Event{}
	var datetime time.Time
	row := pges.db.QueryRow("SELECT id, title, date_time, duration, owner, description FROM events WHERE id::uuid in ($1);", id)

	err := row.Scan(&ev.Uuid, &ev.Title, &datetime, &ev.Duration, &ev.UserId, &ev.Description)
	if err != nil {
		log.Fatal(err)
	}
	ev.Datetime, err = ptypes.TimestampProto(datetime)
	if err != nil {
		log.Fatal(err)
	}
	pges.Log.Info(ev)
	return &ev, nil
}

// UpdateEventByID update event
func (pges *PgEventStorage) UpdateEventByID(ctx context.Context, updateEvent *event.Event) (*event.Event, error) {
	pges.Log.Info("UpdateEventByID")
	dateTime, err := ptypes.Timestamp(updateEvent.Datetime)
	if err != nil {
		log.Fatal(err)
	}
	_, err = pges.db.Exec(`UPDATE events SET title=$1, date_time=$2, duration=$3, owner=$4, description=$5, send=false WHERE id::uuid in ($6)`,
		updateEvent.Title, dateTime, updateEvent.Duration, updateEvent.UserId, updateEvent.Description, updateEvent.Uuid)
	if err != nil {
		log.Fatal(err)
	}
	return updateEvent, nil
}

// DeleteEventByID delete event
func (pges *PgEventStorage) DeleteEventByID(ctx context.Context, id string) error {
	pges.Log.Info("DeleteEventByID")
	_, err := pges.db.Exec(`DELETE FROM events WHERE "id" = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// GetEventList get event list with types
func (pges *PgEventStorage) GetEventList(ctx context.Context, types string, userId string) ([]*event.Event, error) {
	pges.Log.Info("GetEventList")
	var eventList []*event.Event
	var t string
	var datetime time.Time
	var ev event.Event
	switch types {
	case "year":
		t = `SELECT id, title, date_time, duration, owner, description FROM events WHERE (owner = 1 and ((date_time >= now()) and date_time <= now() + make_interval(years := 1)))`
	case "month":
		t = `SELECT id, title, date_time, duration, owner, description FROM events WHERE (owner = 1 and ((date_time >= now()) and date_time <= now() + make_interval(months := 1)))`
	default:
		t = `SELECT id, title, date_time, duration, owner, description FROM events WHERE (owner = 1 and ((date_time >= now()) and date_time <= now() + make_interval(weeks := 1)))`
	}
	row, err := pges.db.Query(t)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for row.Next() {
		err = row.Scan(&ev.Uuid, &ev.Title, &datetime, &ev.Duration, &ev.UserId, &ev.Description)
		if err != nil {
			log.Fatal(err)
		}
		ev.Datetime, err = ptypes.TimestampProto(datetime)
		if err != nil {
			log.Fatal(err)
		}
		eventList = append(eventList, &ev)
	}
	return eventList, nil
}
