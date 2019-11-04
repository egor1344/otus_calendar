package cmd

import (
	"github.com/spf13/viper"
	"log"

	"github.com/egor1344/otus_calendar/calendar/internal/database"
	"github.com/egor1344/otus_calendar/calendar/internal/domain/services"
	grpsServer "github.com/egor1344/otus_calendar/calendar/internal/grps/server"
	"github.com/egor1344/otus_calendar/calendar/pkg/logger"
	"github.com/spf13/cobra"
)

func construct(dsn string) (*grpsServer.CalendarServer, error) {
	eventStorage, err := database.NewPgEventStorage(dsn)
	if err != nil {
		return nil, err
	}
	eventStorage.Log = zapLog
	eventService := &services.Service{
		Database: eventStorage,
		Log:      zapLog,
	}
	server := &grpsServer.CalendarServer{
		EventService: eventService,
		Log:          zapLog,
	}
	return server, nil
}

var host string
var port string
var dsn string

// GrpcServerCmd cobra run server
var GrpcServerCmd = &cobra.Command{
	Use:   "server",
	Short: "run grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		server, err := construct(dsn)
		if err != nil {
			zapLog.Fatal(err)
		}
		address := host + ":" + port
		zapLog.Info(address)
		err = server.RunServer("tcp", address)
		if err != nil {
			zapLog.Fatal(err)
		}
	},
}

func init() {
	//GrpcServerCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Host server")
	GrpcServerCmd.Flags().StringVar(&port, "port", "", "Port server")
	//GrpcServerCmd.Flags().StringVar(&dsn, "dsn", "host=db port=5432 user=postgres dbname=postgres sslmode=disable", "databse connection string")
	l, err := logger.GetLogger()
	zapLog = l
	if err != nil {
		log.Fatal("Error init logger", err)
	}
	err = viper.BindEnv("CALENDAR_PORT")
	if err != nil {
		zapLog.Fatal(err)
	}
	err = viper.BindEnv("DB_DSN")
	if err != nil {
		zapLog.Fatal(err)
	}
	err = viper.BindEnv("CALENDAR_HOST")
	if err != nil {
		zapLog.Fatal(err)
	}
	viper.AutomaticEnv()
	zapLog.Info(viper.AllSettings())
	host = viper.GetString("calendar_host")
	if port == "" {
		port = viper.GetString("calendar_port")
	}
	dsn = viper.GetString("db_dsn")
}
