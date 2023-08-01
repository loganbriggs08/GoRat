package modules

import (
	"github.com/0xAX/notificator"
	"github.com/pterm/pterm"
)

func NewNotification(title string, description string) {
	notificationPusher = notificator.New(notificator.Options{
		DefaultIcon: "icon.png",
		AppName:     "GoRat",
	})

	err := notificationPusher.Push(title, description, "/assets/icon.png", notificator.UR_CRITICAL)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}
}
