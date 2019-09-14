package main

import "fmt"
import (
	"github.com/egor1344/otus_home/calendar/config"
	loggers "github.com/egor1344/otus_home/calendar/logger"
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
}
