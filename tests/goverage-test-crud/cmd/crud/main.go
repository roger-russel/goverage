package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"goverage-test-crud/internal/pkg/api/routers"
	"goverage-test-crud/pkg/database"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var dbConn string = "./database.db"
var dbDriver string = "sqlite3"
var apiPort string = "8080"

func init() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	os.Remove(dbConn)
	db, err := sql.Open(dbDriver, dbConn)
	if err != nil {
		log.WithError(err).Error("Sql open database error")
	}

	defer db.Close()

	sqlDdl := `
	create table users (id integer not null primary key, name text);
	delete from users;
	`

	_, err = db.Exec(sqlDdl)
	if err != nil {
		log.WithError(err).Error(sqlDdl)
		return
	}
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting backend")

	if os.Getenv("API_PORT") != "" {
		apiPort = os.Getenv("API_PORT")
	}

	if os.Getenv("DB_DRIVER") != "" {
		dbDriver = os.Getenv("DB_DRIVER")
	}

	if os.Getenv("DB_CONN") != "" {
		dbConn = os.Getenv("DB_CONN")
	}

	db := database.Connect(dbDriver, dbConn)

	srv := &http.Server{
		Addr:    ":" + apiPort,
		Handler: routers.Get(db),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	os.Remove("./database.db")

	log.Println("Server exiting")
}
