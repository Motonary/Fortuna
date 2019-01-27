package api

import (
	"github.com/motonary/Fortuna/api"
	"net/http"
	"net/http/httptest"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var router *chi.Mux

func TestUnauthorizedRequestHandle(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users", nil)

	api.Router().ServeHTTP(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	t.Logf("responsed status code : %d\n\n", rw.StatusCode)

	if rw.StatusCode != http.StatusUnauthorized {
		t.Fatalf("unexpected status code : %d", rw.StatusCode)
	}
}

func TestAuthorizedRequestHandle(t *testing.T) {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 2})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer "+tokenString)

	t.Logf("Header : %s \n", r.Header)

	api.Router().ServeHTTP(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	t.Logf("responsed status code : %d\n", rw.StatusCode)

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d", rw.StatusCode)
	}
}
