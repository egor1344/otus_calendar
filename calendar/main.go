package main

import (
	"fmt"

	config "github.com/egor1344/otus_calendar/calendar/config"
	server "github.com/egor1344/otus_calendar/calendar/grps/server"
	loggers "github.com/egor1344/otus_calendar/calendar/logger"
	"github.com/spf13/viper"
)

func main() {
	logger, err := loggers.GetLogger()
	if err != nil {
		fmt.Println("Failed initial logger")
		return
	}
	err = config.ReadConfigFile("config", "./config")
	if err != nil {
		logger.Fatal(err)
	}
	host := viper.Get("CALENDAR_HOST")
	port := viper.Get("CALENDAR_PORT")
	address := host + ":" + port
	logger.Infof("address = ", address)
	server.RunServer("tcp", address)
}
