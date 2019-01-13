package api

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"

	"github.com/motonary/Fortuna/entity"
)

var tokenAuth *jwtauth.JWTAuth

type Response struct {
	Status 	int
	User *entity.User
}

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 2})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func Main() {
	addr := ":3000"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, router())
}

func router() http.Handler {
	r := chi.NewRouter()

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })

	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Route("/users", func(r chi.Router) {
		
			// r.Post("/", createUser)
			
			r.Route("/{userID}", func(r chi.Router) {
				r.Get("/", getUser)
				// r.Put("/", updateUser)
				// r.Delete("/", deleteUser)
			})
		})
	})

	return r

}

func getUser(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(401), 401)
		return
	}
	// 以下のエラーが出るのでタイプアサーションの後、intにキャスト
	// panic: interface conversion: interface {} is float64, not int [recovered]
	// panic: interface conversion: interface {} is float64, not int
	// goroutine 36 [running]: -> loop

	userID := int(claims["user_id"].(float64))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("user_id : " + string(userID))

	user, err := dbGetUser(userID)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}

	response := Response{http.StatusOK, user}

	res, err := json.Marshal(response)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func dbGetUser(userID int) (*entity.User, error) {
	return entity.NewUser(userID, "ririco", "ririco@example.com", "test"), nil
}
