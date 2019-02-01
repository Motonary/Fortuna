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

func GetUserBy(name string, data interface{}) (*entity.User, error) {
	var user *entity.User

	switch v := data.(type) {
	case int:
		db.First(&user, name+"=?", v)
	case string:
		db.First(&user, name+"=?", v)
	default:
		log.Printf("undefined type %T!\n", v)
		return nil, errors.New("undefined type of argument")
	}
	dbUserLogger(user)

	return user, nil
}

func UpdateUser(user *entity.User) (*entity.User, error) {
	var userUpdated *entity.User

	userUpdated.Email = user.Email
	db.First(&userUpdated)
	db.Model(&userUpdated).Update(user)
	dbUserLogger(userUpdated)

	return userUpdated, nil
}

func DeleteUser(userID int) error {
	var user *entity.User

	user.ID = userID
	db.First(&user)
	db.Delete(&user)
	dbUserLogger(user)

	return nil
}

func dbUserLogger(user interface{}) {
	data,_ := json.Marshal(user)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(string(data))
}
