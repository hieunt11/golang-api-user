package main

import (
	"log"
	"net/http"
	userapi "user-rancher/cmd/api"
	"user-rancher/cmd/connectdb"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {
	db, err := connectdb.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/user/create", userapi.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}
