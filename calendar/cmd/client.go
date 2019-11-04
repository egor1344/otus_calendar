package cmd

import (
	"context"
	"github.com/egor1344/otus_calendar/calendar/pkg/logger"
	protoServer "github.com/egor1344/otus_calendar/calendar/proto/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

var server string
var title string
var text string
var startTime string
var endTime string

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
		// Неделя
		zapLog.Info("week")
		req := &protoServer.GetEventListRequest{
			Type: protoServer.GetEventListRequest_week,
		}
		resp, err := client.GetEventList(context.Background(), req)
		if err != nil {
			zapLog.Fatal("add event error ", err)
		}
		event := resp.GetEvent()
		zapLog.Info(event)
		// Месяц
		zapLog.Info("month")
		req = &protoServer.GetEventListRequest{
			Type: protoServer.GetEventListRequest_month,
		}
		resp, err = client.GetEventList(context.Background(), req)
		if err != nil {
			zapLog.Fatal("add event error ", err)
		}
		event = resp.GetEvent()
		zapLog.Info(event)
		// Год
		zapLog.Info("year")
		req = &protoServer.GetEventListRequest{
			Type: protoServer.GetEventListRequest_year,
		}
		resp, err = client.GetEventList(context.Background(), req)
		if err != nil {
			zapLog.Fatal("add event error ", err)
		}
		event = resp.GetEvent()
		zapLog.Info(event)
		// Добавление события
		//req := &protoServer.AddEventRequest{
		//	Event: &protoEvent.Event{
		//		Datetime: ptypes.TimestampNow(), Title: "test", Description: "Description", UserId: 1,
		//	},
		//}
		//zapLog.Info(req)
		//resp, err := client.AddEvent(context.Background(), req)
		//if err != nil {
		//	zapLog.Fatal("add event error ", err)
		//}
		//event := resp.GetEvent()
		//zapLog.Info(event)
		//zapLog.Info(event.Uuid)
		//// Получение события
		//reqGet := &protoServer.GetEventRequest{
		//	Id: event.Uuid,
		//}
		//zapLog.Info(reqGet)
		//respGet, err := client.GetEvent(context.Background(), reqGet)
		//if err != nil {
		//	zapLog.Fatal("get event error ", err)
		//}
		//event = respGet.GetEvent()
		//zapLog.Info(event)
		//// Обновление события
		//reqUpdate := &protoServer.UpdateEventRequest{
		//	Event: &protoEvent.Event{
		//		Uuid:        event.Uuid,
		//		Datetime:    ptypes.TimestampNow(),
		//		Title:       "test_update",
		//		Description: "update_description",
		//		UserId:      1,
		//	},
		//}
		//zapLog.Info(reqUpdate)
		//respUpdate, err := client.UpdateEvent(context.Background(), reqUpdate)
		//if err != nil {
		//	zapLog.Fatal("update event error ", err)
		//}
		//event = respUpdate.GetEvent()
		//zapLog.Info(event)
		//// Удаление события
		//reqDelete := &protoServer.DeleteEventRequest{
		//	Id: event.Uuid,
		//}
		//zapLog.Info(reqDelete)
		//respDelete, err := client.DeleteEvent(context.Background(), reqDelete)
		//if err != nil {
		//	zapLog.Fatal("delete event error ", err)
		//}
		//status := respDelete.GetStatus()
		//zapLog.Info(status)
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
