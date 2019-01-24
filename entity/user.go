package entity

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"size:255"`
	Email     string    `json:"email" gorm:"size:255"`
	Password  string    `json:"password" gorm:"size:255"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUser(id int, name string, email string, password string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func UpdateUser(id int) (*User, bool) {
	user := &User{
		ID:       id,
		Name:     "name",
		Email:    "email",
		Password: "password",
	}
	return user, false
}

func DeleteUser(id int) (*User, bool) {
	user := &User{
		ID:       id,
		Name:     "name",
		Email:    "email",
		Password: "password",
	}
	return user, false
}

// func (u *User) IsMale() bool {
// 	return u.Gender == "male"
// }
