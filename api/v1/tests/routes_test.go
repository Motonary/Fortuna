package v1

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"

	api "github.com/motonary/Fortuna/api/v1"
	"github.com/motonary/Fortuna/entity"
)

var (
	router      *chi.Mux
	testUser    *entity.User
	testBody    []byte
	tokenAuth   *jwtauth.JWTAuth
	tokenString string
)

func TestUnauthorizedRequestHandle(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/1", nil)

	api.Router().ServeHTTP(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	t.Logf("responsed status code : %d\n\n", rw.StatusCode)

	if rw.StatusCode != http.StatusUnauthorized {
		t.Fatalf("unexpected status code : %d\n\n", rw.StatusCode)
	}
}

func TestAuthorizeRequestHandle(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/users", bytes.NewBuffer(testBody))

	api.Router().ServeHTTP(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	var responsedBody Response
	responsedByteBody, _ := ioutil.ReadAll(rw.Body)
	json.Unmarshal(responsedByteBody, &responsedBody)
	tokenString = responsedBody.Token

	t.Logf("%v\n", tokenString)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/auth", nil)
	r.Header.Set("Authorization", "Bearer "+tokenString)

	api.Router().ServeHTTP(w, r)
	res := w.Result()
	defer res.Body.Close()

	t.Logf("responsed status code : %d\n\n", rw.StatusCode)

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code : %d\n\n", rw.StatusCode)
	}
}

func setup() {
	println("setup")

	testUser = entity.NewUser(0, "ririco", "ririco@example.com", "password")
	testBody, _ = json.Marshal(testUser)

	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	println(testUser)
}

func teardown() {
	println("teardown")
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}
