package handler

import (
	"os"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	// "github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth
var tokenString string

func TestGetUserHandlerResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer " + tokenString)
	// t.Logf("Header : %s \n", r.Header)
	GetUser(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	t.Logf("responsed status code : %d\n", rw.StatusCode)
	
	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d", rw.StatusCode)
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	t.Logf("responsed json : %s\n", string(b))
}

func TestUpdatetUserHandlerResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer " + tokenString)

  t.Logf("Header : %s \n", r.Header)

	UpdateUser(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	t.Logf("responsed status code : %d\n", rw.StatusCode)
	
	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d", rw.StatusCode)
	}
}

func TestDeleteUserHandlerResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer " + tokenString)

  t.Logf("Header : %s \n", r.Header)

	DeleteUser(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	t.Logf("responsed status code : %d\n", rw.StatusCode)
	
	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d", rw.StatusCode)
	}
}

func setup() {
	println("setup")
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ = tokenAuth.Encode(jwt.MapClaims{"user_id": 3})
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