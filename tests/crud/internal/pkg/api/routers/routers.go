package routers

import (
	"database/sql"
	"net/http"

	"goverage-test-crud/internal/pkg/api/controllers/healthz"
	"goverage-test-crud/internal/pkg/api/controllers/users"

	"github.com/gorilla/mux"
)

//Get Routers
func Get(db *sql.DB) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/users", injector(users.Index, db)).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", injector(users.IndexByID, db)).Methods("GET")
	router.HandleFunc("/users", injector(users.Create, db)).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", injector(users.Delete, db)).Methods("DELETE")
	router.HandleFunc("/healthz", healthz.Index).Methods("GET")

	return router

}

func injector(f func(w http.ResponseWriter, r *http.Request, db *sql.DB), db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, db)
	}
}
