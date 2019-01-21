package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/inflection"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
	Password string `gorm:"size:255"`
}

func main() {
	db := gormConnect()
	defer db.Close()

	db.AutoMigrate(&User{})
	// db.Create(&User{Name: "test", Email: "ririco722@motonary.com", Password: "password"})
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "Motonary"
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := "" // 必要？
	DBNAME := "fortuna"
	OPTION := "charset=utf8&parseTime=True&loc=Local" // enable time.Time

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
