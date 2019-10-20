package main

import (
	"context"
	"errors"
	"github.com/DATA-DOG/godog"
	"github.com/egor1344/otus_calendar/calendar/proto/event"
	"github.com/egor1344/otus_calendar/calendar/proto/server"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"log"
	"os"
)

var dbDsn = os.Getenv("DB_DSN")
var serverPort = os.Getenv("CALENDAR_PORT")
var serverHost = "server"

type serverTest struct {
	//dbDsn string
	DB *sqlx.DB

	Client server.CalendarEventClient

	addEventReq *server.AddEventResponse

	responseStatusCode int
	responseBody       []byte
}

func (test *serverTest) connectDB(interface{}) {
	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	test.DB = db
}

func (test *serverTest) iAddEventGprcrequestTo(arg1 string) (err error) {
	// Добавление события
	conn, err := grpc.Dial(arg1, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error connect to grpc server ", arg1, err)
	}
	defer conn.Close()
	client := server.NewCalendarEventClient(conn)
	test.Client = client
	req := &server.AddEventRequest{
		Event: &event.Event{
			Datetime: ptypes.TimestampNow(), Title: "test", Description: "Description", UserId: 1,
		},
	}
	resp, err := test.Client.AddEvent(context.Background(), req)
	test.addEventReq = resp
	return nil
}

func (test *serverTest) theResponseShouldMatchMyEvent() (err error) {
	resp := test.addEventReq
	ev := resp.GetEvent()
	if ev.Uuid == "" {
		return errors.New("not uuid in event struct")
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	test := new(serverTest)

	s.BeforeScenario(test.connectDB)
	s.Step(`^I AddEvent gprc-request to "([^"]*)"$`, test.iAddEventGprcrequestTo)
	s.Step(`^The response should match my Event$`, test.theResponseShouldMatchMyEvent)

}
