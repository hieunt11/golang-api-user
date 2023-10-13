package usermodels

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Username    string `json: "username"`
	Password    string `json: "password"`
	Description string `json: "description"`
}

var db *sqlx.DB

func CreateUser(user *User) bool {
	if user.Username != "" && user.Password != "" {
		_, err := db.Query("INSERT INTO rancher_user (user_name, password, description) VALUES ('%v', '%v', '%v')", user.Username, user.Password, user.Description)

		if err != nil {
			log.Fatal(err)
			return false
		}

		return true
	} else {
		return false
	}
}
