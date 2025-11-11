package main

import (
	"fmt"

	"github.com/ilonaoken/app/calendar"
	"github.com/ilonaoken/app/events"
)

func main() {
	e, err := events.NewEvent("dentist", "2025-02-12 09:30")
	if err != nil {
		fmt.Println(err)
		return
	}
	calendar.AddEvent("event1", e)
	fmt.Println("Календарь обновлён")
	calendar.ShowEvent()
}
