package database

import (
	"encoding/json"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/inflection"

	"github.com/motonary/Fortuna/entity"
)

func CreateUser(user *entity.User) (*entity.User, error) {
	DB.Create(&user)
	if !DB.NewRecord(user) {
		return user, nil
	}
	return nil, errors.New("User not created")
}

func GetUserBy(name string, data interface{}) (*entity.User, error) {
	user := entity.User{}

	switch v := data.(type) {
	case int:
		DB.First(&user, name+"=?", v)
	case string:
		DB.First(&user, name+"=?", v)
		log.Printf("%v!\n", v)
	default:
		log.Printf("undefined type %T!\n", v)
		return nil, errors.New("undefined type of argument")
	}
	dbUserLogger(&user)

	return &user, nil
}

func UpdateUser(user *entity.User) (*entity.User, error) {
	userUpdated := entity.User{}

	userUpdated.Email = user.Email
	DB.First(&userUpdated)
	DB.Model(&userUpdated).Update(user)
	dbUserLogger(&userUpdated)

	return &userUpdated, nil
}

func DeleteUser(userID int) error {
	user := entity.User{}

	user.ID = userID
	DB.First(&user)
	DB.Delete(&user)
	dbUserLogger(&user)

	return nil
}

func dbUserLogger(user *entity.User) {
	data, _ := json.Marshal(&user)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(string(data))
}
