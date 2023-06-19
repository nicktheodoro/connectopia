package database

import (
	"connectopia-api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Connect opens a connection to the database and returns it
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DatabaseConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
