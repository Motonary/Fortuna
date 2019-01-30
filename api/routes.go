package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var (
	tokenAuth *jwtauth.JWTAuth
	tokenString string
)

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func Main() {
	addr := ":3000/api/v1"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, Router())
}

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
