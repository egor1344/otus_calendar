package main

import (
	"github.com/egor1344/otus_calendar/calendar/cmd"
	"log"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}