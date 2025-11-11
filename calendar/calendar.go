package calendar

import (
	"fmt"

	"github.com/ilonaoken/app/events"
)

var calendar = make(map[string]events.Event)

func AddEvent(key string, e events.Event) {
	calendar[key] = e
	fmt.Println("Событие добавлено:", e.Title)
}
func ShowEvent() {
	for _, value := range calendar {
		fmt.Println(value)
	}
}
