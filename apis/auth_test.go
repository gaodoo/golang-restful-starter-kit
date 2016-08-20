package apis

import (
	"testing"
	"net/http"
)

func TestAuth(t *testing.T) {
	router.Post("/auth", Auth("secret"))
	runAPITests(t, []apiTestCase{
		{"t1 - successful login", "POST", "/auth", `{"username":"demo", "password":"pass"}`, http.StatusOK, ""},
		{"t2 - unsuccessful login", "POST", "/auth", `{"username":"demo", "password":"bad"}`, http.StatusUnauthorized, ""},
	})
}