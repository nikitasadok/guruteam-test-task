package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testTaskGuru/globals"
	"testTaskGuru/models/entities"
	"testTaskGuru/models/requests"
	"testTaskGuru/models/responses"
)

func AddUser (w http.ResponseWriter, r *http.Request) {
	var user entities.User
	var errorResponse responses.ErrorResponse
	var response responses.AddUserResponse
	var request requests.AddUserRequest

	var err error

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		errorResponse.Error = "Error reading your request into byte array"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = json.Unmarshal(body, &request)

	if err != nil {
		errorResponse.Error = "Error unmarshalling the body of your request into a struct"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if !isValidToken(request.Token) {
		errorResponse.Error = "The token is invalid!"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if globals.Users[request.ID] != nil {
		errorResponse.Error = "The user with this ID already exists!"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	user.ID = request.ID
	user.Balance = request.Balance

	globals.Users[user.ID] = &user

	_ = json.NewEncoder(w).Encode(response)
}

func isValidToken(userToken string) bool {
	return userToken == globals.ServerToken
}
