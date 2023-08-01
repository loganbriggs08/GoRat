package modules

import "github.com/0xAX/notificator"

var notificationPusher *notificator.Notificator

func NewNotification() {
	notificationPusher = notificator.New(notificator.Options{
		DefaultIcon: "",
		AppName: "GoRat",
	})
}
