package calendar

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ilonaoken/app/events"
	"github.com/ilonaoken/app/keyboard"
)

var calendar = make(map[string]events.Event)

func AddEvent(e events.Event) {
	calendar[e.Id] = e
	fmt.Println("Событие добавлено:", e.Title)
}

func ShowEvent() {
	for _, value := range calendar {
		fmt.Println(value)
	}
}

func DeleteEvent(key string) {
	delete(calendar, key)
}

func ChangeEvent() {
	fmt.Println("какую задачу вы меняете?")
	deleteOldEvent := bufio.NewScanner(os.Stdin)
	deleteOldEvent.Scan()
	delete(calendar, deleteOldEvent.Text())
	title, date := keyboard.GetEventTitle()
	events.NewEvent(title, date)
}

// func SaveEvents() {

// 	file, err := os.Create(filename)
// 	fmt.Println("error", err)
// 	encoder := gob.NewEncoder(file)
// 	encoder.Encode(tasks)
// 	file.Close()

// }
