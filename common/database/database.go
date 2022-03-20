package database

import (
	"database/sql"
)

func PrepareDatabase() (*sql.DB, error) {
	db, err := preparePostgres()
	if err != nil {
		return nil, err
	}
	return db, nil
}
