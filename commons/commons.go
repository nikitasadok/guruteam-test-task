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

func ProcessJSON (request requests.Request, errorResponse *responses.ErrorResponse, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		errorResponse.Error = "Error reading your request into byte array"
		return err
	}

	print (reflect.TypeOf(request).Name())

	//switch t := reflect.TypeOf(request); t {
	//	case requests.GetUserRequest
	//}



	err = json.Unmarshal(body, &request)


	if err != nil {
		errorResponse.Error = "Error unmarshalling the body of your request into a struct"
		return err
	}

	if !IsValidToken(request.GetToken()) {
		errorResponse.Error = "The token is invalid!"
		return err
	}
	return nil
}

func IsValidToken(userToken string) bool{
	return userToken == globals.ServerToken
}
