package entity

type User struct {
	Id				int			`json:"id"`
	Name 			string	`json:"name"`
	Email			string	`json:"email"`
	Password	string 	`json:"password"`
}

func NewUser(id int ,name string, email string, password string) *User {
	return &User{
					Id: id,
					Name: name,
					Email:  email,
					Password: password,
	}
}

// func (u *User) IsMale() bool {
// 	return u.Gender == "male"
// }
