package windows

import (
	"fmt"

	"github.com/NotKatsu/GoRat/modules/gui/components"

	go_rat "github.com/AllenDang/giu"
)

func MainWindow() {
	go_rat.SingleWindow().Layout(
		go_rat.Label(fmt.Sprint("Logs:")),
		go_rat.ListBox("Logs", []string{"Listening on port 8080...", "New Connection from AABE9B69-76A3-4344-A2CC-9540838DBJAA", "New Connection from AABE9B69-76A3-4344-A2CC-9540838DBJAA"}).Size(950, 200),

		go_rat.Label(fmt.Sprintf("Connected Machines (%d)", 4)),

		go_rat.Table().Columns(
			go_rat.TableColumn("MAC Address"),
			go_rat.TableColumn("OS"),
			go_rat.TableColumn("Name"),
			go_rat.TableColumn("Connected")).Rows(components.CreateClientTable()...))
}
