package healthz

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Index(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
