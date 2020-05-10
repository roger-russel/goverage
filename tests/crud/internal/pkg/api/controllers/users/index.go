package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//IndexByID the user by id
func IndexByID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	log.Info(r)
	vars := mux.Vars(r)
	log.Info("Getting user by id: " + vars["id"])

	id, err := strconv.Atoi(vars["id"])

	db.Begin()
	if err != nil {
		log.WithError(err).Error("Database begin error")
	}

	statement, err := db.Prepare("select id, name from users where id = ?")
	if err != nil {
		log.WithError(err).Error(statement)
	}

	u := User{}
	err = statement.QueryRow(id).Scan(&u.ID, &u.Name)
	if err != nil {
		log.WithError(err).Error(statement)
	}

	output, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

}

//Index the user by id
func Index(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	log.Info(r)
	log.Info("Getting users")

	if _, err := db.Begin(); err != nil {
		log.WithError(err).Error("Database begin error")
	}

	statement, err := db.Prepare("select id, name from users")

	if err != nil {
		log.WithError(err).Error(statement)
	}

	userList := make([]User, 0)
	rows, err := statement.Query()

	for rows.Next() {

		u := User{}

		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			log.Fatal(err)
		}

		userList = append(userList, u)

	}

	if err != nil {
		log.WithError(err).Error(statement)
	}

	output, err := json.Marshal(userList)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)

}
