package server

import (
	"context"
	"log"
	"net"

	db "github.com/egor1344/otus_calendar/calendar/db"
	calendar_server "github.com/egor1344/otus_calendar/calendar/models/server"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) AddEvent(ctx context.Context, in *calendar_server.AddEventRequest) (*calendar_server.AddEventResponse, error) {
	log.Println("add event", in.GetEvent())
	err := db.AddEvent(in.GetEvent())
	if err != nil {
		return nil, err
	}
	response := calendar_server.AddEventResponse{
		Status: "True",
	}
	return &response, nil

}

func (s *server) UpdateEvent(ctx context.Context, in *calendar_server.UpdateEventRequest) (*calendar_server.UpdateEventResponse, error) {
	log.Println("update event")
	response := calendar_server.UpdateEventResponse{
		Status: "True",
	}
	return &response, nil
}
func (s *server) DeleteEvent(ctx context.Context, in *calendar_server.DeleteEventRequest) (*calendar_server.DeleteEventResponse, error) {
	log.Println("delete event")
	response := calendar_server.DeleteEventResponse{
		Status: "True",
	}
	return &response, nil
}

// RunServer - Создание сервера grpc
func RunServer(network, address) error {
	conn, err := net.Listen(network, address)
	log.Println("server run in", address)
	if err != nil {
		log.Fatal(err)
		return err
	}
	gs := grpc.NewServer()
	calendar_server.RegisterCalendarEventServer(gs, &server{})
	err = gs.Serve(conn)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
