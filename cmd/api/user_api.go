package userapi

import (
	"encoding/json"
	"net/http"
	usermodels "user-rancher/cmd/user"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var user usermodels.User

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := usermodels.CreateUser(&user)

		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not create user")
			return
		}

		responseWithJSON(response, http.StatusOK, user)
	}
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
