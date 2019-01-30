package api

import (
	"github.com/motonary/Fortuna/api"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var router *chi.Mux

func TestUnauthorizedRequestHandle(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/1", nil)

	api.Router().ServeHTTP(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	t.Logf("responsed status code : %d\n\n", rw.StatusCode)

	if rw.StatusCode != http.StatusUnauthorized {
		t.Fatalf("unexpected status code : %d\n\n", rw.StatusCode)
	}
}

func TestAuthorizedRequestHandle(t *testing.T) {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 2})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer "+tokenString)

	api.Router().ServeHTTP(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	t.Logf("responsed status code : %d\n\n", rw.StatusCode)

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d\n\n", rw.StatusCode)
	}
}

func setup() {
	println("setup")
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ = tokenAuth.Encode(jwt.MapClaims{"user_id": 1})
}

func teardown() {
	println("teardown")
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}
