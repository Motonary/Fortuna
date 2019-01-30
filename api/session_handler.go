package api

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/motonary/Fortuna/entity"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	user, err := getCreateUserParams(r)
	if err != nil {
		httpErrCheck(w, err, http.StatusInternalServerError)
	}
	user, err = dbGetUserByEmail(user)
	if err != nil {
		httpErrCheck(w, err, http.StatusInternalServerError)
	}
	_, tokenString, _ = tokenAuth.Encode(jwt.MapClaims{"user_id": user.ID})

	response := Response{http.StatusOK, user, tokenString}
	jsonResponse(w, response)
}

func dbGetUserByEmail(user *entity.User) (*entity.User, error) {
	return user, nil
}
