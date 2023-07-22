package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type ConnectionData struct {
	ID            string `json:"id"`
	LastHeartbeat string `json:"last_heartbeat"`
}

func CreateTables() bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return false
	} else {
		_, err := database.Exec("CREATE TABLE IF NOT EXISTS connections(id VARCHAR(100), last_heartbeat_time VARCHAR(255))")

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
		_, err := database.Exec("INSERT INTO connections(id, last_heartbeat_time) VALUES (?,?)", ID, time.Now())

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func GetConnectionData(ID string) (ConnectionData, bool) {

}
