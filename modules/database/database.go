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

func CreateTables() bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return false
	} else {
		_, err := database.Exec(`CREATE TABLE IF NOT EXISTS connections(id VARCHAR(100), last_heartbeat_time VARCHAR(255), connection_time VARCHAR(255))`)

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
		log.Fatal(err)
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
