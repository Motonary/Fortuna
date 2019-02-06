package database_test

import (
	"testing"

	db "github.com/motonary/Fortuna/database"
)

func TestGetUserByName(t *testing.T) {
	user, err := db.GetUserBy("name", "ririco")
	if err != nil {
		t.Fatalf("unexpected error occured: %v\n", err)
	}

	if user.Name != "ririco" {
		t.Fatalf("user name not matched\n")
	}
}

func TestGetUserByEmail(t *testing.T) {
	user, err := db.GetUserBy("email", "ririco@example.com")
	if err != nil {
		t.Fatalf("unexpected error occured: %v\n", err)
	}

	if user.Name != "ririco" {
		t.Fatalf("user name not matched\n")
	}
}

func TestGetUserByNotFound(t *testing.T) {
	user, err := db.GetUserBy("name", "nouser")
	if err != nil {
		t.Fatalf("unexpected error occured: %v\n", err)
	}

	if user.ID != 0 {
		t.Fatalf("should be nil but somthing was returned: %v\n", user)
	}
}
