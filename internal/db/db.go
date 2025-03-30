package db

import (
	"database/sql"

	"github.com/nyakovchuk/vsch_church_bot/config"
)

func ConnectDB(config *config.Config) (*sql.DB, error) {
	dbDriver := getDriver(config.DbDriver)

	db, err := sql.Open(dbDriver, config.Dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDriver(dbDriver string) string {
	switch dbDriver {
	case "sqlite":
		return "sqlite3"
	default:
		return dbDriver
	}
}
