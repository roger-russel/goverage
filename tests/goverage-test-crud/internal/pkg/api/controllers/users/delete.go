package users

import (
	"database/sql"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

//Delete User
func Delete(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	log.Info(r)
	log.Info("Deleting user")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	transaction, err := db.Begin()
	if err != nil {
		log.WithError(err).Error("Database begin error")
	}

	statement := "delete from users where id = $1"
	_, err = transaction.Exec(statement, id)
	if err != nil {
		log.WithError(err).Error(statement)
	}
	log.Infof("Created user '%d'.", id)

	transaction.Commit()
}
