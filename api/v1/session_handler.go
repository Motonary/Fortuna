package v1

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	db "github.com/motonary/Fortuna/database"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	user, err := getUserParams(r)
	if err != nil {
		httpErrCheck(w, err, http.StatusInternalServerError)
	}
	user, err = db.GetUserBy("email", user.Email)
	if err != nil {
		httpErrCheck(w, err, http.StatusInternalServerError)
	}
	_, tokenString, _ = tokenAuth.Encode(jwt.MapClaims{"user_id": user.ID})

	response := Response{http.StatusOK, user, tokenString}
	jsonResponse(w, response)
}
