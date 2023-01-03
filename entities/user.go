package entities

import "fmt"

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user User) ToString() string {
	return fmt.Sprint("id: %s\nName: %s\nPassword: %n\n", user.Id, user.Username, user.Password)
}
