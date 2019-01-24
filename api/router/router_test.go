package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

func TestUnauthorizedRequestHandle(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/1", nil)

	router.ServeHTTP(w, r)
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

	router := router()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer "+tokenString)

	t.Logf("Header : %s \n", r.Header)

	router.ServeHTTP(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	t.Logf("responsed status code : %d\n", rw.StatusCode)

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d", rw.StatusCode)
	}
}
