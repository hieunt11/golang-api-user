package apis

import (
	"api-user/entities"
	"api-user/models"
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(response http.ResponseWriter, statusCode int, mgs string) {
	responseWithJSON(response, statusCode, map[string]string{
		"error": mgs,
	})
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func FindUser(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 0 {
		responseWithError(response, http.StatusBadRequest, "URL Param id is missing")
		return
	}
	user, err := models.GetUserId(ids[0])
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
		return
	}
	responseWithJSON(response, http.StatusOK, user)
}

func GetListUser(response http.ResponseWriter, request *http.Request) {
	user := models.GetAllUser()
	responseWithJSON(response, http.StatusOK, user)
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var user entities.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
		log.Print(err.Error())
	} else {
		result := models.CreateUser(&user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not create user")
			return
		}
		responseWithJSON(response, http.StatusOK, "Create user successfully")
	}
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	var user *entities.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
		log.Print(err.Error())
	} else {
		result := models.UpdateUser(user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not update user")
			return
		}
		responseWithJSON(response, http.StatusOK, "Update user successfully")
	}
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 0 {
		responseWithError(response, http.StatusBadRequest, "URL Param id is missing")
		return
	}
	result := models.DeleteUser(ids[0])
	if !result {
		responseWithError(response, http.StatusBadRequest, "Could not delete user")
		return
	}
	responseWithJSON(response, http.StatusOK, "Delete user successfully")
}
