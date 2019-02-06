package v1

import (
	"github.com/motonary/Fortuna/entity"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var (
	tokenAuth   *jwtauth.JWTAuth
	tokenString string
)

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func Main() {
	// TODO: productionとdevで区別
	addr := ":3000/api/v1"
	log.Printf("Starting server on %v\n", addr)
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

		r.Get("/auth", authRequest)

		r.Route("/users/{userID}", func(r chi.Router) {
			r.Get("/", GetUser)
			r.Put("/", UpdateUser)
			r.Delete("/", DeleteUser)
		})
	})
	return mux
}

func issueTokenString(token *jwtauth.JWTAuth, user *entity.User) string {
	claims := jwt.MapClaims{
		"admin":   false,
		"iat":     time.Now(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": user.ID}
	_, tokenString, _ := tokenAuth.Encode(claims)

	return tokenString
}

func authRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("in AuthRequest")
	userID, _, err := getAuthorizedUserParams(r)
	if err != nil {
		httpErrCheck(w, err, http.StatusUnauthorized)
	}
	log.Printf("user_id : %d\n", userID)
	response := Response{http.StatusOK, nil, ""}
	jsonResponse(w, response)
}

func httpErrCheck(w http.ResponseWriter, err error, statusCode int) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(err)
	http.Error(w, http.StatusText(statusCode), statusCode)
}
