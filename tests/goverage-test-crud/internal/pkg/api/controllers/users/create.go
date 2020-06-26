package users

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Create a new User
func Create(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	log.Info(r)
	log.Info("Creating user")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var u User

	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	transaction, err := db.Begin()
	if err != nil {
		log.WithError(err).Error("Database begin error")
	}

	statement := "insert into users (id, name) values ($1, $2)"
	_, err = transaction.Exec(statement, u.ID, u.Name)
	if err != nil {
		log.WithError(err).Error(statement)
		http.Error(w, err.Error(), 500)
	}
	log.Infof("Created user '%d'.", u.ID)

	output, err := json.Marshal(u)
	if err != nil {
		log.WithError(err).Error(output)
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	transaction.Commit()
}
