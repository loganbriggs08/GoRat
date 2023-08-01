package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pterm/pterm"
)

type ConnectionData struct {
	ID                string
	LastHeartbeatTime string
	ConnectionTime    string
}

type EventData struct {
	Recipient string
	EventType string
	Extra     string
}

func CreateTables() bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return false
	} else {
		_, err := database.Exec(`
                                 CREATE TABLE IF NOT EXISTS connections(id VARCHAR(255), last_heartbeat_time VARCHAR(255), connection_time VARCHAR(255));
								 CREATE TABLE IF NOT EXISTS events(recipient VARCHAR(255), type VARCHAR(100), extra VARCHAR(500));
							     CREATE TABLE IF NOT EXISTS event_responses(sender VARCHAR(255), response VARCHAR(255))`)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func ConnectionNew(ID string) bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return false
	} else {
		_, err := database.Exec("INSERT INTO connections(id, last_heartbeat_time, connection_time) VALUES (?,?,?)", ID, time.Now(), time.Now())

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func GetConnectionData(ID string) string {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, last_heartbeat_time FROM connections WHERE id = ?", ID)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var connectionData ConnectionData

		err := rows.Scan(&connectionData.ID, &connectionData.LastHeartbeatTime)

		if err != nil {
			pterm.Fatal.WithFatal(true).Println(err)
		} else {
			return connectionData.LastHeartbeatTime
		}

	}
	return ""
}

func DeleteConnection(ID string) bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return false
	} else {
		_, err := database.Exec("DELETE FROM connections WHERE id=?", ID)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func UpdateConnection(ID string) bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return false
	} else {
		_, err := database.Exec("UPDATE connections SET last_heartbeat_time=? WHERE id=?", time.Now(), ID)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func GetConnections() []ConnectionData {
	var Connections []ConnectionData

	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return Connections
	} else {
		rows, err := database.Query("SELECT id, last_heartbeat_time, connection_time FROM connections")

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var connectionData ConnectionData

			err := rows.Scan(&connectionData.ID, &connectionData.LastHeartbeatTime, &connectionData.ConnectionTime)

			if err != nil {
				pterm.Fatal.WithFatal(true).Println(err)
			} else {
				Connections = append(Connections, connectionData)
			}
		}

		return Connections
	}
}

func GetConnectionTime(ID string) string {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, last_heartbeat_time, connection_time FROM connections WHERE id = ?", ID)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var connectionData ConnectionData

		err := rows.Scan(&connectionData.ID, &connectionData.LastHeartbeatTime, &connectionData.ConnectionTime)

		if err != nil {
			pterm.Fatal.WithFatal(true).Println(err)
		} else {
			return connectionData.ConnectionTime
		}

	}
	return ""
}

func DeleteClientEvent(ID string, event_type string, extra string) bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
		return false
	}

	_, err = database.Exec("DELETE FROM events WHERE recipient = ? AND type = ? AND extra = ?", ID, event_type, extra)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
		return false
	} else {
		return true
	}
}

func GetClientEvent(ID string) EventData {
	var EventDataResult EventData
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
		return EventData{Recipient: "None", EventType: "None", Extra: "None"}
	}
	defer database.Close()

	rows, err := database.Query("SELECT recipient, type, extra FROM events WHERE recipient = ? LIMIT 1", ID)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)

		return EventData{Recipient: "None", EventType: "None", Extra: "None"}
	}
	defer rows.Close()

	for rows.Next() {
		var eventDataCurrent EventData

		err := rows.Scan(&eventDataCurrent.Recipient, &eventDataCurrent.EventType, &eventDataCurrent.Extra)

		if err != nil {
			pterm.Fatal.WithFatal(true).Println(err)
		}
		return eventDataCurrent
	}

	return EventDataResult
}

func CreateNewClientEvent(recipient string, eventType string, extra string) bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}

	_, err = database.Exec("INSERT INTO events(recipient, type, extra) VALUES (?, ?, ?)", recipient, eventType, extra)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
		return false
	} else {
		return true
	}
}
