package database

import (
	"database/sql"
	"os"
)

var currentDrive string
var currentStrConn string
var currentDB *sql.DB

//Connect to database
func Connect(drive string, strConn string) *sql.DB {
	var err error

	currentDrive = drive
	currentStrConn = strConn

	if currentDB, err = sql.Open(drive, strConn); err != nil {
		panic(err)
	}

	return currentDB

}

func Close() {
	if err := currentDB.Close(); err != nil {
		panic(err)
	}
}

func Remove() {
	os.Remove(currentStrConn)
}
