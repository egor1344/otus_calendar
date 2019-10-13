package cmd

import (
	"context"
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
	Use:   "grpc_client",
	Short: "Run grpc client",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(server)
		conn, err := grpc.Dial(server, grpc.WithInsecure())
		if err != nil {
			log.Fatal("Error connect to grpc server ", server, err)
		}
		defer conn.Close()
		client := protoServer.NewCalendarEventClient(conn)
		req := &protoServer.AddEventRequest{
			Event: &protoEvent.Event{
				Datetime: ptypes.TimestampNow(), Title: "test", Description: "Description", UserId: 1,
			},
		}
		log.Println(req)
		resp, err := client.AddEvent(context.Background(), req)
		if err != nil {
			log.Fatal("add event error ", err)
		}
		log.Println(resp.GetResult())
		// req = &protoServer.AddEventRequest{
		// 	Event: &protoEvent.Event{
		// 		Date: ptypes.TimestampNow(), Title: "test", Description: "Description", UserId: 1,
		// 	},
		// }
		// log.Println(req)
		// resp, err = client.AddEvent(context.Background(), req)
		// if err != nil {
		// 	log.Fatal("add event error ", err)
		// }
	},
}

func init() {
	GrpcClientCmd.Flags().StringVar(&server, "server", "server:8000", "host server")
	GrpcClientCmd.Flags().StringVar(&title, "title", "", "event title")
	GrpcClientCmd.Flags().StringVar(&text, "text", "", "event text")
	GrpcClientCmd.Flags().StringVar(&startTime, "start-time", "", "event startTime")
	GrpcClientCmd.Flags().StringVar(&endTime, "end-time", "", "event endTime")

}
