package database

import "os"

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := gormConnect()
	defer db.Close()
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "fortuna"
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := "" // 必要？
	DBNAME := "fortuna"
	OPTION := "charset=utf8&parseTime=True&loc=Local" // enable time.Time

	CONNECT = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
