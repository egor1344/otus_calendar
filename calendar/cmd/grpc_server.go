package cmd

import (
	"log"

	"github.com/egor1344/otus_calendar/calendar/internal/database"
	"github.com/egor1344/otus_calendar/calendar/internals/domain/services"
	"github.com/egor1344/otus_calendar/calendar/internals/grpc/server"
	"github.com/spf13/cobra"
)

func construct(dsn string) (*server.CalendarServer, error) {
	eventStorage, err := database.NewPgEventStorage(dsn)
	if err != nil {
		return nil, err
	}
	eventService := &services.EventService{
		EventStorage: eventStorage,
	}
	server := &server.CalendarServer{
		EventService: eventService,
	}
	return server, nil
}

var host string
var port string
var dsn string

// GrpcServerCmd cobra run server
var GrpcServerCmd = &cobra.Command{
	Use:   "grpc_server",
	Short: "run grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		server, err := construct()
		if err != nil {
			log.Fatal(err)
		}
		address := host + ":" + port
		err = server.Serve(address)
	},
}

func init() {
	GrpcServerCmd.Flags().StringVar(&host, "host", "localhost", "Host server")
	GrpcServerCmd.Flags().StringVar(&port, "port", "8000", "Port server")
	GrpcServerCmd.Flags().StringVar(&dsn, "dsn", "host=127.0.0.1 user=event_user password=event_pwd dbname=event_db", "databse connection string")
}
