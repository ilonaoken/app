package chooseAction

import (
	"fmt"

	"github.com/ilonaoken/app/calendar"
	"github.com/ilonaoken/app/events"
)

func ChooseAction(act string, title string, date string) {
	if act == "add" {
		e, err := events.NewEvent(title, date)
		if err != nil {
			fmt.Println(err)
			return
		}
		calendar.AddEvent(e)
		fmt.Println("Календарь обновлён")
	} else if act == "delete" {
		calendar.DeleteEvent(title)
	} else if act == "change" {
		calendar.ChangeEvent()
	} else if act == "print" {
		calendar.ShowEvent()
	}

}
