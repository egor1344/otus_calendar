package db

import (
	"log"

	"github.com/golang/protobuf/ptypes"

	event "otus_home/calendar/models/event"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func initConnetction() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=marko dbname=calendar password=g sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return db, nil
}

// AddEvent Добавление события в БД
func AddEvent(event *event.Event) error {
	db, err := initConnetction()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	tx := db.MustBegin()
	log.Println(event.Date, ptypes.TimestampString(event.Date))
	tx.MustExec("INSERT INTO events (title, owner, duration, description, date_time, before_duration) VALUES ($1, $2, $3, $4, $5, $6)", event.Title, event.UserId, event.Duration, event.Description, ptypes.TimestampString(event.Date), event.BeforeDuration)
	tx.Commit()
	return nil
}

func main() {
	db, err := sqlx.Connect("postgres", "user=marko dbname=calendar password=g sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	_ = db
}
