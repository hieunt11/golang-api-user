package models

import (
	"api-user/connect"
	"api-user/entities"
	"errors"
	"log"
)

var listUser []entities.User

func GetAllUser() []entities.User {
	db := connect.Openconnect()
	results, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Print(err.Error())
	}
	for results.Next() {
		var user entities.User
		err = results.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			log.Print(err.Error())
		}
		//fmt.Printf("ID: %s, Username: %s, Password: %s\n", user.Id, user.Username, user.Password)
		listUser = append(listUser, user)
	}
	db.Close()
	return listUser
}

func GetUserId(id string) (entities.User, error) {
	db := connect.Openconnect()
	var user entities.User
	// scan single row
	err := db.QueryRow("SELECT * FROM user WHERE id=?", id).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		log.Print(err.Error())
		return user, errors.New("User not found")
	} else {
		return user, nil
	}
}

func DeleteUser(id string) bool {
	db := connect.Openconnect()
	listUser := GetAllUser()
	for _, user := range listUser {
		if user.Id == id {
			_, err := db.Query("DELETE FROM user WHERE id=?", id)
			if err != nil {
				log.Print(err.Error())
				db.Close()
				return false
			} else {
				db.Close()
				return true
			}
		}
	}
	return false
}

func CreateUser(user *entities.User) bool {
	db := connect.Openconnect()
	_, err := db.Query("INSERT INTO user (id, username, password) VALUES (?, ?, ?)", user.Id, user.Username, user.Password)
	if err != nil {
		print(err.Error())
		db.Close()
		return false
	} else {
		db.Close()
		return true
	}
}

func UpdateUser(updateuser *entities.User) bool {
	db := connect.Openconnect()
	listUser := GetAllUser()
	for _, user := range listUser {
		if user.Id == updateuser.Id {
			_, err := db.Query("UPDATE user SET username = ?, password = ? WHERE id = ?", updateuser.Username, updateuser.Password, updateuser.Id)
			if err != nil {
				log.Print(err.Error())
				db.Close()
				return false
			} else {
				db.Close()
				return true
			}
		}
	}
	return false
}
