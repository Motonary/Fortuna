package database_test

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/motonary/Fortuna/entity"
	test "github.com/motonary/Fortuna/testdata"
)

var (
	DB *gorm.DB
)

func TestBuildTestData(t *testing.T) {
	user := entity.User{}
	DB.First(&user, "name=?", "ririco")

	if user.Name != "ririco" {
		t.Fatalf("record not found\n\n")
	}
}

func setup() {
	println("setup")
	DB = test.BuildDB()
}

func teardown() {
	println("teardown")
	test.CleanDB(DB)
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}
