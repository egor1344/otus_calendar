package server

import (
	"context"
	"log"
	"net"

	"github.com/egor1344/otus_calendar/calendar/internal/domain/services"
	calendar_server "github.com/egor1344/otus_calendar/calendar/proto/server"

	"google.golang.org/grpc"
)

// CalendarServer - Реализует работу с grpc сервером
type CalendarServer struct {
	EventService *services.Service
}

// AddEvent add event
func (s *CalendarServer) AddEvent(ctx context.Context, in *calendar_server.AddEventRequest) (*calendar_server.AddEventResponse, error) {
	log.Println("add event", in.GetEvent())
	event := in.GetEvent()
	log.Println(s)
	newEvent, err := s.EventService.AddEvent(ctx, event.GetTitle(), event.GetDate(), event.GetDuration(), event.GetDescription(), event.GetUserId())
	log.Println(newEvent)
	if err != nil {
		response := &calendar_server.AddEventResponse{
			Result: &calendar_server.AddEventResponse_Error{
				Error: "error happend",
			},
		}
		return response, err
	}
	response := &calendar_server.AddEventResponse{
		Result: &calendar_server.AddEventResponse_Event{
			Event: event,
		},
	}
	return response, nil

}

// GetEvent get event
func (s *CalendarServer) GetEvent(ctx context.Context, in *calendar_server.GetEventRequest) (*calendar_server.GetEventResponse, error) {
	log.Println("get event")
	response := &calendar_server.GetEventResponse{
		Result: &calendar_server.GetEventResponse_Event{
			Event: nil,
		},
	}
	return response, nil
}

// UpdateEvent udpate event
func (s *CalendarServer) UpdateEvent(ctx context.Context, in *calendar_server.UpdateEventRequest) (*calendar_server.UpdateEventResponse, error) {
	log.Println("update event")
	response := &calendar_server.UpdateEventResponse{
		Result: &calendar_server.UpdateEventResponse_Event{
			Event: nil,
		},
	}
	return response, nil
}

// DeleteEvent delete event
func (s *CalendarServer) DeleteEvent(ctx context.Context, in *calendar_server.DeleteEventRequest) (*calendar_server.DeleteEventResponse, error) {
	log.Println("delete event")
	response := calendar_server.DeleteEventResponse{
		Status: "True",
	}
	return &response, nil
}

// RunServer - Создание сервера grpc
func (s *CalendarServer) RunServer(network, address string) error {
	conn, err := net.Listen(network, address)
	log.Println("server run in", network, address)
	if err != nil {
		log.Fatal(err)
		return err
	}
	gs := grpc.NewServer()
	calendar_server.RegisterCalendarEventServer(gs, s)
	err = gs.Serve(conn)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
