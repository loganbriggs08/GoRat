package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTables() bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return false
	} else {
		_, err := database.Exec("CREATE TABLE IF NOT EXISTS connections(id VARCHAR(100), mac_address VARCHAR(100), operating_system VARCHAR(100), computer_name VARCHAR(100))")

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func ConnectionNew(ID string, MACAddress string, OS string, Name string) bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return false
	} else {
		_, err := database.Exec("INSERT INTO connections(id, mac_address, operating_system, computer_name) VALUES (?,?,?,?)", ID, MACAddress, OS, Name)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}
