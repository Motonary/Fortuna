package api

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"

	handler "github.com/motonary/Fortuna/api/handler"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 2})
	log.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func Main() {
	addr := ":3000"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, router())
}

func router() http.Handler {
	mux := chi.NewRouter()

	mux.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Route("/users", func(r chi.Router) {
		
			r.Post("/", handler.CreateUser)
			
			r.Route("/{userID}", func(r chi.Router) {
				r.Get("/", handler.GetUser)
				r.Put("/", handler.UpdateUser)
				r.Delete("/", handler.DeleteUser)
			})
		})
	})

	return mux

}
