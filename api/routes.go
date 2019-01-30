package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func Router() http.Handler {
	mux := chi.NewRouter()

	// Authorization not required
	mux.Group(func(r chi.Router) {
		r.Post("/session", CreateSession)
		r.Post("/users", CreateUser)
	})

	// JWT Authorization required
	mux.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Route("/users/{userID}", func(r chi.Router) {
			r.Get("/", GetUser)
			r.Put("/", UpdateUser)
			r.Delete("/", DeleteUser)
		})
	})
	return mux
}

func httpErrCheck(w http.ResponseWriter, err error, statusCode int) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(err)
	http.Error(w, http.StatusText(statusCode), statusCode)
}
