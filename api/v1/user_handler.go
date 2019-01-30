package v1

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"

	"github.com/motonary/Fortuna/entity"
)

type Response struct {
	Status int          `json:"status"`
	User   *entity.User `json:"user,omitempty"`
	Token  string       `json:"token,omitempty"`
}

// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := getCreateUserParams(r)
	if err != nil {
		httpErrCheck(w, err, http.StatusInternalServerError)
	}

	err = dbCreateUser(user)
	if err != nil {
		httpErrCheck(w, err, http.StatusInternalServerError)
	}
	_, tokenString, _ = tokenAuth.Encode(jwt.MapClaims{"user_id": user.ID})

	response := Response{http.StatusOK, user, tokenString}
	jsonResponse(w, response)
}

// GET /users/:id
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := getAuthorizedUserParams(r)
	if err != nil {
		httpErrCheck(w, err, http.StatusUnauthorized)
	}

	user, err := dbGetUser(userID)
	if err != nil {
		httpErrCheck(w, err, http.StatusBadRequest)
	}
	
	response := Response{http.StatusOK, user, ""}
	jsonResponse(w, response)
}

// PUT /users/:id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := getAuthorizedUserParams(r)
	if err != nil {
		httpErrCheck(w, err, http.StatusUnauthorized)
	}

	user, err := dbUpdateUser(userID)
	if err != nil {
		httpErrCheck(w, err, http.StatusInternalServerError)
	}

	response := Response{http.StatusOK, user, ""}
	jsonResponse(w, response)
}

// DELETE /users/:id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := getAuthorizedUserParams(r)
	if err != nil {
		httpErrCheck(w, err, http.StatusUnauthorized)
	}

	err = dbDeleteUser(userID)
	if err != nil {
		httpErrCheck(w, err, http.StatusInternalServerError)
	}

	response := Response{http.StatusOK, nil, ""}
	jsonResponse(w, response)
}

func dbCreateUser(user *entity.User) error {
	entity.NewUser(user.ID, user.Name, user.Email, user.Password)
	return nil
}

func dbGetUser(userID int) (*entity.User, error) {
	return entity.NewUser(userID, "ririco", "ririco@example.com", "test"), nil
}

func dbUpdateUser(userID int) (*entity.User, error) {
	user, err := entity.UpdateUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func dbDeleteUser(userID int) error {
	_, err := entity.DeleteUser(userID)
	if err != nil {
		return err
	}
	return nil
}

func getCreateUserParams(r *http.Request) (*entity.User, error) {
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

func getAuthorizedUserParams(r *http.Request) (int, error) {
	var userID int

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return 0, err
	}
	// 以下のエラーが出るのでタイプアサーションの後、intにキャスト
	// panic: interface conversion: interface {} is float64, not int [recovered]
	// goroutine 36 [running]: -> loop
	if claims["user_id"] == nil {
		return 0, errors.New("JWT claims not included")
	} else {
		userID = int(claims["user_id"].(float64))
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("user_id : %d\n", userID)

	return userID, nil
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
