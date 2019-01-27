package api

import (
	"fmt"
	"log"
	"net/http"

	// jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"

	"github.com/motonary/Fortuna/api/session"
)

var tokenAuth *jwtauth.JWTAuth
var globalSessions *session.Manager

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	// _, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 2})
	log.Printf("DEBUG: a sample jwt is %v\n\n", tokenAuth)

	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
}

func Main() {
	addr := ":3000/v1"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, Router())
}

func Router() http.Handler {
	mux := chi.NewRouter()

	// Authorization not required
	mux.Group(func(r chi.Router) {
		r.Get("/", rootHandler)
		r.Post("/session", CreateSession)
		r.Post("/users", CreateUser)
	})

	// JWT Authorization required
	mux.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Route("/users", func(r chi.Router) {
			r.Route("/{userID}", func(r chi.Router) {
				r.Get("/", GetUser)
				r.Put("/", UpdateUser)
				r.Delete("/", DeleteUser)
			})
		})
	})
	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	
}

// func login(w http.ResponseWriter, r *http.Request) {
// 	sess := globalSessions.SessionStart(w, r)
// 	r.ParseForm()
// 	if r.Method == "GET" {
// 		t, _ := template.ParseFiles("login.gtpl")
// 		w.Header().Set("Content-Type", "text/html")
// 		t.Execute(w, sess.Get("username"))
// 	} else {
// 		sess.Set("username", r.Form["username"])
// 		http.Redirect(w, r, "/", 302)
// 	}
// }
