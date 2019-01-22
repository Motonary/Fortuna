package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth"

	"github.com/motonary/Fortuna/entity"
)

type Response struct {
	Status int          `json:"status"`
	User   *entity.User `json:"user", omitempty`
}

// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {

}

// GET /users/:id
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, errCode := getUserParams(r)
	if errCode != http.StatusOK {
		http.Error(w, http.StatusText(errCode), errCode)
	}

	user, err := dbGetUser(userID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	response := Response{http.StatusOK, user}
	jsonResponse(w, response)
}

// PUT /users/:id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, errCode := getUserParams(r)
	if errCode != http.StatusOK {
		http.Error(w, http.StatusText(errCode), errCode)
	}
	user, errCode := dbUpdateUser(userID)
	if errCode != http.StatusOK {
		http.Error(w, http.StatusText(errCode), errCode)
	}

	response := Response{http.StatusOK, user}
	jsonResponse(w, response)
}

// DELETE /users/:id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, errCode := getUserParams(r)
	if errCode != http.StatusOK {
		http.Error(w, http.StatusText(errCode), errCode)
	}
	user, errCode := dbDeleteUser(userID)
	if errCode != http.StatusOK {
		http.Error(w, http.StatusText(errCode), errCode)
	}

	response := Response{http.StatusOK, user}
	jsonResponse(w, response)
}

func dbGetUser(userID uint) (*entity.User, error) {
	return entity.NewUser(userID, "ririco", "ririco@example.com", "test"), nil
}

func dbUpdateUser(userID uint) (*entity.User, int) {
	var errCode int
	_, err := entity.UpdateUser(userID)
	if err {
		errCode = http.StatusInternalServerError
	}
	return nil, errCode
}

func dbDeleteUser(userID uint) (*entity.User, int) {
	var errCode int
	_, err := entity.DeleteUser(userID)
	if err {
		errCode = http.StatusInternalServerError
	}
	return nil, errCode
}

func getUserParams(r *http.Request) (uint, int) {
	var userID uint

	_, claims, err := jwtauth.FromContext(r.Context())
	statusCode := http.StatusOK
	if err != nil {
		statusCode = http.StatusUnauthorized
	}
	// 以下のエラーが出るのでタイプアサーションの後、intにキャスト
	// panic: interface conversion: interface {} is float64, not int [recovered]
	// goroutine 36 [running]: -> loop
	if claims["user_id"] == nil {
		statusCode = http.StatusInternalServerError
	} else {
		userID = uint(claims["user_id"].(float64))
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("user_id : " + string(userID))

	return userID, statusCode
}

func jsonResponse(w http.ResponseWriter, response Response) {
	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
