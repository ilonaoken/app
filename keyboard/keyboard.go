package keyboard

import (
	"bufio"
	"fmt"
	"os"
)

func GetEvent() (string, string, string) {
	actionScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите необходимое действие:add,delete,change,print,close...")
	actionScanner.Scan()
	action := actionScanner.Text()
	if action == "add" {
		title, startAt := GetEventTitle()
		return title, startAt, action
	}

	return "", "", action
}
func GetEventTitle() (string, string) {
	titleScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите название события:")
	titleScanner.Scan()
	title := titleScanner.Text()

	fmt.Println("Введите дату события:")
	dateScanner := bufio.NewScanner(os.Stdin)
	dateScanner.Scan()
	startAt := dateScanner.Text()
	return title, startAt
}
