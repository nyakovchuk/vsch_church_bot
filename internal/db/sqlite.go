package db

import (
	"database/sql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
