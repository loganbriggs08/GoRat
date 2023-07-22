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
		database.Exec("CREATE TABLE IF NOT EXISTS connections()")
	}
}
