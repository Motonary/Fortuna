package database

import (
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/inflection"
	yaml "gopkg.in/yaml.v2"
)

var (
	DB   *gorm.DB
	seed map[string]map[string]string
)

func init() {
	yml := loadConfig(getConfigFile())
	_ = yaml.Unmarshal(yml, &seed)

	DB = Connect()
	log.Printf("database connected\n")
	log.Printf("%s\n", os.Getenv("GO_ENV"))
}

func Connect() *gorm.DB {
	conf := seed[os.Getenv("GO_ENV")]

	CONNECT := conf["user"] + ":" + conf["pass"] + "@" + conf["protocool"] + "/" + conf["db"] + "?" + conf["option"]
	DB, err := gorm.Open(conf["driver"], CONNECT)

	if err != nil {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		log.Print(err)
	}

	return DB
}

func getConfigFile() string {
	if os.Getenv("GO_ENV") == "circleci" {
		return "../../../config/database.ci.yml"
	}
	if _, err := os.Stat("config/database.yml"); err == nil {
		return "config/database.yml"
	}
	if _, err := os.Stat("../config/database.yml"); err == nil {
		return "../config/database.yml"
	}
	return "../../config/database.yml"
}

func loadConfig(file_path string) []byte {
	yml, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		log.Println(err)
	}
	return yml
}

func dbLogger() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println()
}
