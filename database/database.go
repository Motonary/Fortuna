package database

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/inflection"
)

func main() {
	db, err := gormConnect()
	defer db.Close()

	if err != nil {
		return
	}
}

func gormConnect() (*gorm.DB, error) {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := ""
	DBNAME := os.Getenv("DB_NAME")
	OPTION := "charset=utf8&parseTime=True&loc=Local" // enable time.Time

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		log.Print(err)
	}

	return db, err
}
