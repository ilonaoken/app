package main

import (
	"github.com/ilonaoken/app/chooseAction"
	"github.com/ilonaoken/app/keyboard"
)

func main() {
	for {
		title, date, action := keyboard.GetEvent()
		if action == "close" {
			break
		}
		chooseAction.ChooseAction(action, title, date)

	}
}
