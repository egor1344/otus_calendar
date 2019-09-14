package main

import (
	"context"
	"log"

	event "github.com/egor1344/otus_calendar/calendar/models/event"
	calendar_server "github.com/egor1344/otus_calendar/calendar/models/server"

	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc"
)

func main() {
	port := "127.0.0.1:3000"
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("Error connect to grpc server: %v\n", err)
	}
	client := calendar_server.NewCalendarEventClient(conn)
	request := &calendar_server.AddEventRequest{
		Event: &event.Event{Id: 1, Date: ptypes.TimestampNow(), Title: "test", Description: "Description"},
	}
	response, err := client.AddEvent(context.Background(), request)
	if err != nil {
		log.Fatalf("Error got response: %v\n", err)
	}
	log.Printf("Got gRPC response: %v\n", response.Status)
}
