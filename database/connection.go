package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Connect to database
func OpenConnection(path string) error {
	var err error
	DB, err = sql.Open("sqlite3", path)
	return err
}

// Close database
func CloseConnection() error {
	return DB.Close()
}
