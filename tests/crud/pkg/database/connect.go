package database

import (
	"database/sql"
)

//Connect to database
func Connect(drive string, strConn string) (db *sql.DB) {
	var err error

	if db, err = sql.Open(drive, strConn); err != nil {
		panic(err)
	}

	return db

}
