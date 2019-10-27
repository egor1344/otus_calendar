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

var dbDsn = os.Getenv("TEST_DB_DSN")
var serverPort = os.Getenv("CALENDAR_PORT")
var serverHost = "server"

type serverTest struct {
	//dbDsn string
	DB *sqlx.DB

	Client server.CalendarEventClient

	addEventReq    *server.AddEventResponse
	getEventReq    *server.GetEventResponse
	updateEventReq *server.UpdateEventResponse

	getEventListReq *server.GetEventListResponse

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

func (test *serverTest) flushDB(interface{}) {
	_, err := test.DB.Exec("truncate table events;")
	if err != nil {
		log.Fatal(err)
	}
}

func (test *serverTest) addEvents(interface{}) {
	// week test
	test.DB.Exec("INSERT INTO public.events (date_time, description, duration, owner, title, before_time_pull, id, send) VALUES (now() + make_interval(days := 1), 'd1', 0, 1, 't1', 0, gen_random_uuid(), false);")
	test.DB.Exec("INSERT INTO public.events (date_time, description, duration, owner, title, before_time_pull, id, send) VALUES (now() + make_interval(days := 3), 'd2', 0, 1, 't2', 0, gen_random_uuid(), false);")
	test.DB.Exec("INSERT INTO public.events (date_time, description, duration, owner, title, before_time_pull, id, send) VALUES (now() + make_interval(days := 4), 'd3', 0, 1, 't3', 0, gen_random_uuid(), false);")
	// month test
	test.DB.Exec("INSERT INTO public.events (date_time, description, duration, owner, title, before_time_pull, id, send) VALUES (now() + make_interval(months :=1, days := 1), 'd4', 0, 1, 't4', 0, gen_random_uuid(), false);")
	test.DB.Exec("INSERT INTO public.events (date_time, description, duration, owner, title, before_time_pull, id, send) VALUES (now() + make_interval(months :=1, days := 3), 'd5', 0, 2, 't5', 0, gen_random_uuid(), false);")
	test.DB.Exec("INSERT INTO public.events (date_time, description, duration, owner, title, before_time_pull, id, send) VALUES (now() + make_interval(months :=1, days := 4), 'd6', 0, 3, 't6', 0, gen_random_uuid(), false);")
	test.DB.Exec("INSERT INTO public.events (date_time, description, duration, owner, title, before_time_pull, id, send) VALUES (now() + make_interval(days := 14), 'd7', 0, 3, 't7', 0, gen_random_uuid(), false);")

	// month test
	test.DB.Exec("INSERT INTO public.events (date_time, description, duration, owner, title, before_time_pull, id, send) VALUES (now() + make_interval(months :=2, days := 1), 'd8', 0, 1, 't8', 0, gen_random_uuid(), false);")
	test.DB.Exec("INSERT INTO public.events (date_time, description, duration, owner, title, before_time_pull, id, send) VALUES (now() + make_interval(months :=3, days := 3), 'd9', 0, 2, 't9', 0, gen_random_uuid(), false);")
	test.DB.Exec("INSERT INTO public.events (date_time, description, duration, owner, title, before_time_pull, id, send) VALUES (now() + make_interval(months :=4, days := 4), 'd10', 0, 3, 't10', 0, gen_random_uuid(), false);")
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
	if err != nil {
		return err
	}
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

func (test *serverTest) iGetEventGprcrequestTo(arg1 string) error {
	// Получение события
	conn, err := grpc.Dial(arg1, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error connect to server", arg1, err)
	}
	defer conn.Close()
	client := server.NewCalendarEventClient(conn)
	test.Client = client
	req := &server.GetEventRequest{Id: test.addEventReq.GetEvent().Uuid}
	resp, err := test.Client.GetEvent(context.Background(), req)
	if err != nil {
		return err
	}
	test.getEventReq = resp
	return nil
}

func (test *serverTest) theResponseMustContainMyEvent() error {
	resp := test.getEventReq
	ev := resp.GetEvent()
	if ev.Uuid == "" {
		return errors.New("not uuid in event struct")
	}
	if ev.Title != test.addEventReq.GetEvent().Title {
		return errors.New("return another event")
	}
	return nil
}

func (test *serverTest) iUpdateEventGprcrequestTo(arg1 string) error {
	// Обновление события
	conn, err := grpc.Dial(arg1, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error connect to server", arg1, err)
	}
	defer conn.Close()
	client := server.NewCalendarEventClient(conn)
	test.Client = client
	req := &server.UpdateEventRequest{Event: &event.Event{
		Uuid: test.addEventReq.GetEvent().Uuid, Datetime: ptypes.TimestampNow(),
		Title: "update_test", Description: "update_description", UserId: 1,
	}}
	resp, err := test.Client.UpdateEvent(context.Background(), req)
	if err != nil {
		return err
	}
	test.updateEventReq = resp
	return nil
}

func (test *serverTest) theResponseMustContainMyUpdateEvent() error {
	ev := test.updateEventReq.GetEvent()
	if ev.Title != "update_test" {
		return errors.New("Update title not equal old title")
	}
	if ev.Uuid != test.addEventReq.GetEvent().Uuid {
		return errors.New("return another event")
	}
	return nil
}

func (test *serverTest) iGetEventListWithTypeWeekGprcrequestTo(arg1 string) error {
	// Получение событий за неделю
	conn, err := grpc.Dial(arg1, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error connect to server", arg1, err)
	}
	defer conn.Close()
	client := server.NewCalendarEventClient(conn)
	test.Client = client
	req := &server.GetEventListRequest{Type: server.GetEventListRequest_week}
	resp, err := test.Client.GetEventList(context.Background(), req)
	if err != nil {
		return err
	}
	test.getEventListReq = resp
	return nil
}

func (test *serverTest) theResponseMustContainEventListOnWeek(arg1 string) error {
	// Получение событий за неделю
	log.Println(test.getEventListReq.Event)
	return nil
}

func FeatureContext(s *godog.Suite) {
	test := new(serverTest)

	s.BeforeScenario(test.connectDB)
	// add event
	s.Step(`^I AddEvent gprc-request to "([^"]*)"$`, test.iAddEventGprcrequestTo)
	s.Step(`^The response should match my Event$`, test.theResponseShouldMatchMyEvent)
	// get event
	s.Step(`^I Get event gprc-request to "([^"]*)"$`, test.iGetEventGprcrequestTo)
	s.Step(`^The response must contain my event$`, test.theResponseMustContainMyEvent)
	// update event
	s.Step(`^I UpdateEvent gprc-request to "([^"]*)"$`, test.iUpdateEventGprcrequestTo)
	s.Step(`^The response must contain my update event$`, test.theResponseMustContainMyUpdateEvent)
	// get list event
	s.BeforeScenario(test.flushDB)
	s.BeforeScenario(test.addEvents)
	// week
	s.Step(`^I GetEventList with type week gprc-request to "([^"]*)"$`, test.iGetEventListWithTypeWeekGprcrequestTo)
	s.Step(`^The response must contain event list on week$`, test.theResponseMustContainEventListOnWeek)
	// month
	//s.Step(`^I GetEventList with type month gprc-request to "([^"]*)"$`, iGetEventListWithTypeMonthGprcrequestTo)
	//s.Step(`^The response must contain event list on month$`, theResponseMustContainEventListOnMonth)
	// year
	//s.Step(`^I GetEventList with type year gprc-request to "([^"]*)"$`, iGetEventListWithTypeYearGprcrequestTo)
	//s.Step(`^The response must contain event list on year$`, theResponseMustContainEventListOnYear)

}
