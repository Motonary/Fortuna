package api

import (
	"bytes"
	"encoding/json"
	// "io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	
	"github.com/motonary/Fortuna/api"
	"github.com/motonary/Fortuna/entity"
)

var tokenAuth *jwtauth.JWTAuth
var tokenString string

func TestCreateUserHandler(t *testing.T) {
	userParams := entity.NewUser(0, "ririco", "ririco@example.com", "test")
	body, _ := json.Marshal(userParams)
	
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/users",  bytes.NewBuffer(body))

	api.Router().ServeHTTP(w, r)

	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d", rw.StatusCode)
	}
}

func TestGetUserHandlerResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer "+tokenString)

	api.Router().ServeHTTP(w, r)

	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d", rw.StatusCode)
	}
}

func TestUpdatetUserHandlerResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer "+tokenString)

	api.Router().ServeHTTP(w, r)

	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d", rw.StatusCode)
	}
}

func TestDeleteUserHandlerResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer "+tokenString)

	api.Router().ServeHTTP(w, r)

	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d", rw.StatusCode)
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
