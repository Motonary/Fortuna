package v1

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	
	api "github.com/motonary/Fortuna/api/v1"
	"github.com/motonary/Fortuna/entity"
)

func TestCreateSessionHandler(t *testing.T) {
	user := entity.NewUser(0, "", "ririco@example.com", "test")
	body, _ := json.Marshal(user)
	
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/session",  bytes.NewBuffer(body))

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
