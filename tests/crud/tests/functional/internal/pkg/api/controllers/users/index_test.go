package users

import (
	"goverage-test-crud/internal/pkg/api/routers"
	"goverage-test-crud/pkg/database"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {

	defer database.Remove()

	req, _ := http.NewRequest("GET", "/users", nil)
	res := httptest.NewRecorder()

	db := database.Connect("sqlite3", "./database.db")
	routers.Get(db).ServeHTTP(res, req)

	if res.Body.String() != "[]" {
		t.Error("Expected \"[]\" but received: ", res.Body.String())
	}

}
