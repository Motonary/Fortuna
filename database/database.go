package database

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/inflection"

	"github.com/motonary/Fortuna/entity"
)

func main() {
	db, err := gormConnect()
	defer db.Close()

	if err != nil {
		return
	}

	db.AutoMigrate(&entity.User{})
}

func gormConnect() (*gorm.DB, error) {
	DBMS := "mysql"
	USER := "Motonary"
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := ""
	DBNAME := "fortuna"
	OPTION := "charset=utf8&parseTime=True&loc=Local" // enable time.Time

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		log.Print(err)
	}

	return db, err
}
