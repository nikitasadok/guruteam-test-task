package commons

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testTaskGuru/globals"
	"testTaskGuru/models/requests"
	"testTaskGuru/models/responses"
)

func ProcessJSON (request requests.Request, r *http.Request) (responses.ErrorResponse, error, requests.Request) {
	var errorResponse responses.ErrorResponse
	var concreteRequest requests.Request

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		errorResponse.Error = "Error reading your request into byte array"
		return errorResponse, err, nil
	}

	print (reflect.TypeOf(request).Name())

	switch t:= reflect.TypeOf(request).Name(); t {
	case "AddUserRequest":
		concreteRequest = request.(requests.AddUserRequest)
	case "GetUserRequest":
		concreteRequest = request.(requests.GetUserRequest)
	case "AddDepositRequest":
		concreteRequest = request.(requests.AddDepositRequest)
	case "TransactionRequest":
		concreteRequest = request.(requests.TransactionRequest)
	}

	err = json.Unmarshal(body, &concreteRequest)

	if err != nil {
		errorResponse.Error = "Error unmarshalling the body of your request into a struct"
		return errorResponse, err, nil
	}

	if !IsValidToken(request.GetToken()) {
		errorResponse.Error = "The token is invalid!"
		return errorResponse, err, nil
	}
	return errorResponse, nil, concreteRequest
}

func IsValidToken(userToken string) bool{
	return userToken == globals.ServerToken
}
