package api

import (
	"log"
	"net/http"

	"github.com/go-chi/jwtauth"
)

var (
	tokenAuth *jwtauth.JWTAuth
	tokenString string
)

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func main() {
	// TODO: 環境によって分岐させる
	addr := ":3000/api/v1"
	log.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, Router())
}
