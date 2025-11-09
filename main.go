package main

import (
	"fmt"
	"time"

	"github.com/ilonaoken/app/calendar"
	"github.com/ilonaoken/app/events"
)

func main() {
	e := events.Event{
		Title:   "Встреча",
		StartAt: time.Now(),
	}
	calendar.AddEvent("event1", e)
	fmt.Println("Календарь обновлён")
}
