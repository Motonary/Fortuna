package v1

import (
	"net/http"

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

	tokenString = issueTokenString(tokenAuth, user)

	response := Response{http.StatusOK, user, tokenString}
	jsonResponse(w, response)
}
