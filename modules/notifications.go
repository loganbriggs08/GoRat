package modules

import "github.com/gen2brain/beeep"

func NewNotification(title string, description string) {
	err := beeep.Notify(title, description, "/assets/icon.png")
	if err != nil {
		panic(err)
	}
}
