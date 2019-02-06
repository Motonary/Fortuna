package v1

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	api "github.com/motonary/Fortuna/api/v1"
	"github.com/motonary/Fortuna/entity"
)

type Response struct {
	Status int          `json:"status"`
	User   *entity.User `json:"user,omitempty"`
	Token  string       `json:"token,omitempty"`
}

func TestCreateUserHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/users", bytes.NewBuffer(testBody))

	api.Router().ServeHTTP(w, r)

	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d\n\n", rw.StatusCode)
	}

	var response Response
	bytes, _ := ioutil.ReadAll(rw.Body)
	json.Unmarshal(bytes, &response)

	if response.User.Name != testUser.Name {
		t.Fatalf("response data is unexpected : %v\n\n", response.User)
	}
}

func TestGetUserHandlerResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer "+tokenString)

	api.Router().ServeHTTP(w, r)

	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d\n\n", rw.StatusCode)
	}
}

// updateのサーバとクライアントのインターフェースが未定のためペンディング
// func TestUpdatetUserHandlerResponse(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest("PUT", "/users/2", nil)
// 	r.Header.Set("Authorization", "Bearer "+tokenString)

// 	api.Router().ServeHTTP(w, r)

// 	rw := w.Result()
// 	defer rw.Body.Close()

// 	if rw.StatusCode != http.StatusOK {
// 		t.Fatalf("unexpected status code : %d\n\n", rw.StatusCode)
// 	}
// }

func TestDeleteUserHandlerResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer "+tokenString)

	api.Router().ServeHTTP(w, r)

	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d\n\n", rw.StatusCode)
	}
}
