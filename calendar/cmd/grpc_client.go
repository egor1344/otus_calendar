package cmd

import (
	"context"
	"log"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"time"

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
	Use: "grpc_client",
	Short: "Run grpc client",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(server, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		client := api.NewCalendarServiceClient(conn)
		st, err := parseTs(startTime)
		if err != nil {
			log.Fatal(err)
		}
		et, err := parseTs(endTime)
		if err != nil {
			log.Fatal(err)
		}
		req := &api.CreateEventRequest {
			Title: title, 
			Text: text,
			StartTime: st,
			EndTime: et,
		}
		resp, err := client.CreateEvent(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.GetError() != ""{
			log.Fatal(resp.GetError())
		} else {
			log.Println(resp.GetEnvent().id)
		}
	},
}

func init(){
	GrpcClientCmd.Flags().StringVar(&server, "server", "localhost:8000", "host server")
	GrpcClientCmd.Flags().StringVar(&title, "title", "", "event title")
	GrpcClientCmd.Flags().StringVar(&text, "text", "", "event text")
	GrpcClientCmd.Flags().StringVar(&startTime, "start-time", "", "event startTime")
	GrpcClientCmd.Flags().StringVar(&endTime, "end-time", "", "event endTime")
	
	
}