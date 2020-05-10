package healthz

import (
	"goverage-test-crud/internal/pkg/api/routers"
	"goverage-test-crud/pkg/database"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestIndex(t *testing.T) {

	defer database.Remove()

	req, _ := http.NewRequest("GET", "/healthz", nil)
	res := httptest.NewRecorder()

	db := database.Connect("sqlite3", "./database.db")
	routers.Get(db).ServeHTTP(res, req)

	if res.Body.String() != "ok" {
		t.Error("Expected \"ok\" but received: ", res.Body.String())
	}

}
