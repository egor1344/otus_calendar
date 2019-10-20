package server

import (
	"context"
	"go.uber.org/zap"
	"log"
	"net"

	"github.com/egor1344/otus_calendar/calendar/internal/domain/services"
	calendar_server "github.com/egor1344/otus_calendar/calendar/proto/server"

	"google.golang.org/grpc"
)

// CalendarServer - Реализует работу с grpc сервером
type CalendarServer struct {
	EventService *services.Service
	Log          *zap.SugaredLogger
}

// AddEvent add event
func (s *CalendarServer) AddEvent(ctx context.Context, in *calendar_server.AddEventRequest) (*calendar_server.AddEventResponse, error) {
	s.Log.Info("add event", in.GetEvent())
	event := in.GetEvent()
	s.Log.Info(s)
	newEvent, err := s.EventService.AddEvent(ctx, event.GetTitle(), event.GetDatetime(), event.GetDuration(), event.GetDescription(), event.GetUserId(), event.GetBeforeTimePull())
	s.Log.Info(newEvent)
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
			Event: newEvent,
		},
	}
	return response, nil

}

// GetEvent get event
func (s *CalendarServer) GetEvent(ctx context.Context, in *calendar_server.GetEventRequest) (*calendar_server.GetEventResponse, error) {
	s.Log.Info("get event")
	uuid := in.GetId()
	s.Log.Info(s)
	event, err := s.EventService.GetEvent(ctx, uuid)
	s.Log.Info(event)
	if err != nil {
		response := &calendar_server.GetEventResponse{
			Result: &calendar_server.GetEventResponse_Error{
				Error: "error happend",
			},
		}
		return response, err
	}
	response := &calendar_server.GetEventResponse{
		Result: &calendar_server.GetEventResponse_Event{
			Event: event,
		},
	}
	return response, nil
}

// UpdateEvent udpate event
func (s *CalendarServer) UpdateEvent(ctx context.Context, in *calendar_server.UpdateEventRequest) (*calendar_server.UpdateEventResponse, error) {
	s.Log.Info("update event")
	response := &calendar_server.UpdateEventResponse{
		Result: &calendar_server.UpdateEventResponse_Event{
			Event: nil,
		},
	}
	return response, nil
}

// DeleteEvent delete event
func (s *CalendarServer) DeleteEvent(ctx context.Context, in *calendar_server.DeleteEventRequest) (*calendar_server.DeleteEventResponse, error) {
	s.Log.Info("delete event")
	response := calendar_server.DeleteEventResponse{
		Status: "True",
	}
	return &response, nil
}

// RunServer - Создание сервера grpc
func (s *CalendarServer) RunServer(network, address string) error {
	conn, err := net.Listen(network, address)
	s.Log.Info("server run in", network, address)
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
