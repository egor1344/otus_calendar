package cmd

import (
	"log"

	"github.com/egor1344/otus_calendar/calendar/internal/database"
	"github.com/egor1344/otus_calendar/calendar/internal/domain/services"
	grpsServer "github.com/egor1344/otus_calendar/calendar/internal/grps/server"
	"github.com/spf13/cobra"
)

func construct(dsn string) (*grpsServer.CalendarServer, error) {
	eventStorage, err := database.NewPgEventStorage(dsn)
	if err != nil {
		return nil, err
	}
	eventService := &services.Service{
		Database: eventStorage,
	}
	server := &grpsServer.CalendarServer{
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
		server, err := construct(dsn)
		if err != nil {
			log.Fatal(err)
		}
		address := host + ":" + port
		err = server.RunServer("tcp", address)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	GrpcServerCmd.Flags().StringVar(&host, "host", "localhost", "Host server")
	GrpcServerCmd.Flags().StringVar(&port, "port", "8000", "Port server")
	GrpcServerCmd.Flags().StringVar(&dsn, "dsn", "host=127.0.0.1 user=event_user password=event_pwd dbname=event_db", "databse connection string")
}
