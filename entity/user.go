package entity

import "time"

type User struct {
	ID        int       `json:"id,omitempty" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"size:255"`
	Email     string    `json:"email" gorm:"size:255"`
	Password  string    `json:"-" gorm:"size:255"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func NewUser(id int, name string, email string, password string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func UpdateUser(id int) (*User, error) {
	user := &User{
		ID:       id,
		Name:     "name",
		Email:    "email",
		Password: "password",
	}
	return user, nil
}

func DeleteUser(id int) (*User, error) {
	user := &User{
		ID:       id,
		Name:     "name",
		Email:    "email",
		Password: "password",
	}
	return user, nil
}
