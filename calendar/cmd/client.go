package cmd

import (
	"context"
	"github.com/egor1344/otus_calendar/calendar/pkg/logger"
	"github.com/spf13/viper"
	"log"
	"time"

	protoEvent "github.com/egor1344/otus_calendar/calendar/proto/event"
	protoServer "github.com/egor1344/otus_calendar/calendar/proto/server"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var server string
var title string
var text string
var startTime string
var endTime string

const tsLayout = "2006-01-02T15:04:05"

func parseTs(s string) (*timestamp.Timestamp, error) {
	t, err := time.Parse(tsLayout, s)
	if err != nil {
		return nil, err
	}
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

// GrpcClientCmd cobra run client
var GrpcClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run grpc client",
	Run: func(cmd *cobra.Command, args []string) {
		zapLog.Info(server)
		if server == "" {
			server = "server:" + viper.GetString("calendar_port")
		}
		zapLog.Info(server)
		conn, err := grpc.Dial(server, grpc.WithInsecure())
		if err != nil {
			zapLog.Fatal("Error connect to grpc server ", server, err)
		}
		defer conn.Close()
		client := protoServer.NewCalendarEventClient(conn)
		req := &protoServer.AddEventRequest{
			Event: &protoEvent.Event{
				Datetime: ptypes.TimestampNow(), Title: "test", Description: "Description", UserId: 1,
			},
		}
		zapLog.Info(req)
		resp, err := client.AddEvent(context.Background(), req)
		if err != nil {
			zapLog.Fatal("add event error ", err)
		}
		event := resp.GetEvent()
		zapLog.Info(event)
		zapLog.Info(event.Uuid)
		reqGet := &protoServer.GetEventRequest{
			Id: event.Uuid,
		}
		zapLog.Info(reqGet)
		respGet, err := client.GetEvent(context.Background(), reqGet)
		if err != nil {
			zapLog.Fatal("get event error ", err)
		}
		event = respGet.GetEvent()
		zapLog.Info(event)
	},
}

func init() {
	GrpcClientCmd.Flags().StringVar(&server, "server", "", "host server")
	GrpcClientCmd.Flags().StringVar(&title, "title", "", "event title")
	GrpcClientCmd.Flags().StringVar(&text, "text", "", "event text")
	GrpcClientCmd.Flags().StringVar(&startTime, "start-time", "", "event startTime")
	GrpcClientCmd.Flags().StringVar(&endTime, "end-time", "", "event endTime")
	l, err := logger.GetLogger()
	zapLog = l
	if err != nil {
		log.Fatal("Error init logger", err)
	}
	err = viper.BindEnv("CALENDAR_PORT")
	if err != nil {
		zapLog.Fatal(err)
	}
	viper.AutomaticEnv()
	zapLog.Info(viper.AllSettings())
}
