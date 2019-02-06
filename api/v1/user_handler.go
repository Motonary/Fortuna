package v1

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth"

	db "github.com/motonary/Fortuna/database"
	"github.com/motonary/Fortuna/entity"
)

type Response struct {
	Status int          `json:"status"`
	User   *entity.User `json:"user,omitempty"`
	Token  string       `json:"token,omitempty"`
}

// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserParams(r)
	if err != nil {
		httpErrCheck(w, err, http.StatusInternalServerError)
	}

	createdUser, dbErr := db.CreateUser(user)
	if dbErr != nil {
		httpErrCheck(w, dbErr, http.StatusInternalServerError)
	}
	tokenString = issueTokenString(tokenAuth, createdUser)

	response := Response{http.StatusOK, createdUser, tokenString}
	jsonResponse(w, response)
}

// GET /users/:id
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, _, err := getAuthorizedUserParams(r)
	if err != nil {
		httpErrCheck(w, err, http.StatusUnauthorized)
	}

	user, err := db.GetUserBy("id", userID)
	if err != nil {
		httpErrCheck(w, err, http.StatusBadRequest)
	}

	response := Response{http.StatusOK, user, ""}
	jsonResponse(w, response)
}

// PUT /users/:id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	_, userParams, err := getAuthorizedUserParams(r)
	if err != nil {
		httpErrCheck(w, err, http.StatusUnauthorized)
	}

	user, err := db.UpdateUser(userParams)
	if err != nil {
		httpErrCheck(w, err, http.StatusInternalServerError)
	}

	response := Response{http.StatusOK, user, ""}
	jsonResponse(w, response)
}

// DELETE /users/:id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, _, err := getAuthorizedUserParams(r)
	if err != nil {
		httpErrCheck(w, err, http.StatusUnauthorized)
	}

	err = db.DeleteUser(userID)
	if err != nil {
		httpErrCheck(w, err, http.StatusInternalServerError)
	}

	response := Response{http.StatusOK, nil, ""}
	jsonResponse(w, response)
}

func getUserParams(r *http.Request) (*entity.User, error) {
	var user *entity.User
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("response : %v\n", string(body))

	return user, nil
}

func getAuthorizedUserParams(r *http.Request) (int, *entity.User, error) {
	var userID int

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return 0, nil, err
	}
	// 以下のエラーが出るのでタイプアサーションの後、intにキャスト
	// panic: interface conversion: interface {} is float64, not int [recovered]
	// goroutine 36 [running]: -> loop
	if claims["user_id"] == nil {
		return 0, nil, errors.New("JWT claims not included")
	} else {
		userID = int(claims["user_id"].(float64))
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("user_id : %d\n", userID)

	user, _ := getUserParams(r)

	return userID, user, nil
}

func jsonResponse(w http.ResponseWriter, response Response) {
	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("response : %s\n", string(res))

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
