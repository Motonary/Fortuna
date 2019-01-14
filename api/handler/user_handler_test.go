package handler

import (
	"os"
	"net/http"
	"net/http/httptest"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth
var tokenString string
var mux *chi.Mux

func TestGetUserHandlerResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer " + tokenString)

	mux.ServeHTTP(w,r)
	
	GetUser(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d", rw.StatusCode)
	}
}

func TestUpdatetUserHandlerResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer " + tokenString)

	mux.ServeHTTP(w,r)

	UpdateUser(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d", rw.StatusCode)
	}
}

func TestDeleteUserHandlerResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer " + tokenString)

	mux.ServeHTTP(w,r)

	DeleteUser(w, r)
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

	mux = chi.NewRouter()

	mux.Group(func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator)
		
			r.Post("/", CreateUser)
			
			r.Route("/{userID}", func(r chi.Router) {
				r.Get("/", GetUser)
				r.Put("/", UpdateUser)
				r.Delete("/", DeleteUser)
			})
		})
	})
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