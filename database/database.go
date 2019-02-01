package database

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/inflection"
	yaml "gopkg.in/yaml.v2"

	"github.com/motonary/Fortuna/entity"
)

var (
	db *gorm.DB
)

func Migrate() {
	for _, model := range []interface{}{
		&entity.User{},
	} {
		if err := db.AutoMigrate(model).Error; err != nil {
			log.Println(err)
		} else {
			log.Println("Auto migrating", reflect.TypeOf(model).Name(), "...")
		}
	}
}

func Connect() {
	yml, err := ioutil.ReadFile("config/database.yml")
	if err != nil {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		log.Println(err)
	}

	t := make(map[interface{}]interface{})
  _ = yaml.Unmarshal([]byte(yml), &t)

	cnf := t[os.Getenv("FORTUNAENV")].(map[interface {}]string)

	CONNECT := cnf["user"] + ":" + cnf["pass"] + "@" + cnf["protocol"] + "/" + cnf["db"] + "?" + cnf["option"]
	db, err = gorm.Open(t["driver"].(string), CONNECT)

	if err != nil {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		log.Print(err)
	}
}

func dbLogger() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println()
}
