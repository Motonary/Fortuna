package testdata

import (
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	yaml "gopkg.in/yaml.v2"

	"github.com/motonary/Fortuna/database"
	"github.com/motonary/Fortuna/entity"
)

type Data struct {
  Users []entity.User `yaml:"users"`
}

var (
	DB *gorm.DB
	seed Data
)

func init() {
	yml := loadConfig(getConfigFile())
	_ = yaml.Unmarshal(yml, &seed)

	DB = database.Connect()
}

func BuildDB() *gorm.DB {
	err := runMakeSeeds(DB, seed)
	if err != nil {
		panic("making seeds couldn't run through\n")
	}
	
	return DB
}

func CleanDB(db *gorm.DB) {
	err := runDestroySeeds(db, seed)
	if err != nil {
		panic("destorying seeds couldn't run through\n")
	}
}

func getConfigFile() string {
	if _, err := os.Stat("config/testdatas.yml"); err == nil {
		return "config/testdatas.yml"
	}
	return "../config/testdatas.yml"
}

func loadConfig(file_path string) []byte {
	yml, err := ioutil.ReadFile(file_path)
	if err != nil {
		panic("couldn't read testdata config\n")
	}
	return yml
}

func runMakeSeeds(db *gorm.DB, seed Data) error {
	for _, user := range seed.Users {
		db.Create(&user)
	}

	return nil
}

func runDestroySeeds(db *gorm.DB, seed Data) error {
	for _, user := range seed.Users {
		db.Delete(&user)
	}

	return nil
}
