package main

import (
	"api-user/apis"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user/find", apis.FindUser).Methods("GET")
	router.HandleFunc("/api/v1/user/getall", apis.GetListUser).Methods("GET")
	router.HandleFunc("/api/v1/user/create", apis.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/update", apis.UpdateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/delete", apis.DeleteUser).Methods("DELETE")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		panic(err.Error())
	}
}
