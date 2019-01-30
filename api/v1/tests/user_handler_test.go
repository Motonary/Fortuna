package v1

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/jwtauth"
	
	api "github.com/motonary/Fortuna/api/v1"
	"github.com/motonary/Fortuna/entity"
)

var tokenAuth *jwtauth.JWTAuth
var tokenString string

type Response struct {
	Status int          `json:"status"`
	User   *entity.User `json:"user,omitempty"`
	Token  string       `json:"token,omitempty"`
}

func TestCreateUserHandler(t *testing.T) {
	user := entity.NewUser(1, "ririco", "ririco@example.com", "test")
	body, _ := json.Marshal(user)
	
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/users",  bytes.NewBuffer(body))

	api.Router().ServeHTTP(w, r)

	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d\n\n", rw.StatusCode)
	}

	bytes,_ := ioutil.ReadAll(rw.Body)
	if strings.Contains(string(body), string(bytes)) {
		t.Fatalf("response data is unexpected : %s\n\n", string(bytes))
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

func TestUpdatetUserHandlerResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/users/2", nil)
	r.Header.Set("Authorization", "Bearer "+tokenString)

	api.Router().ServeHTTP(w, r)

	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d\n\n", rw.StatusCode)
	}
}

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
