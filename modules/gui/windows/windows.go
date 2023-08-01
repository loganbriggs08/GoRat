package windows

import (
	"fmt"
	"time"

	"github.com/NotKatsu/GoRat/modules/database"
	"github.com/NotKatsu/GoRat/modules/gui/components"
	"github.com/pterm/pterm"

	go_rat "github.com/AllenDang/giu"
)

func connectedMachinesCount() int64 {
	var connectionCount int64
	connections := database.GetConnections()

	for _, conn := range connections {
		storedTime, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", conn.LastHeartbeatTime)

		if err != nil {
			pterm.Fatal.WithFatal(true).Println("There was an error when the database tables were trying to be created.")
		}
		currentTime := time.Now()
		timeDifference := currentTime.Sub(storedTime)

		if timeDifference.Seconds() < 5 {
			connectionCount += 1
		}
	}

	return connectionCount
}
func MainWindow() {
	go_rat.SingleWindow().Layout(

		go_rat.Label(fmt.Sprint("Connection Logs:")),
		go_rat.ListBox("Logs", components.CreateClientArray()).Size(go_rat.Auto, 200),

		go_rat.Label(fmt.Sprintf("Connected Machines (%d)", connectedMachinesCount())),

		go_rat.Table().Columns(
			go_rat.TableColumn("MAC Address"),
			go_rat.TableColumn("OS"),
			go_rat.TableColumn("Name"),
			go_rat.TableColumn("Connected")).Rows(components.CreateClientTable()...))
}
