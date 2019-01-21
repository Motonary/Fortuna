package entity

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"size:255"`
	Email     string    `json:"email" gorm:"size:255"`
	Password  string    `json:"password" gorm:"size:255"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUser(id uint, name string, email string, password string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func UpdateUser(id uint) (*User, bool) {
	user := &User{
		ID:       id,
		Name:     "name",
		Email:    "email",
		Password: "password",
	}
	return user, false
}

func DeleteUser(id uint) (*User, bool) {
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
