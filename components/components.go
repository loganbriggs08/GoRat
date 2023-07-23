package components

import (
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