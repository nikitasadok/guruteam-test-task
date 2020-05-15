package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testTaskGuru/globals"
	"testTaskGuru/models/requests"
	"testTaskGuru/models/responses"
)

func GetUser (w http.ResponseWriter, r *http.Request) {
	var request requests.GetUserRequest
	var response responses.GetUserResponse
	var errorResponse responses.ErrorResponse

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

	if globals.Users[request.ID] == nil {
		errorResponse.Error = "There is no such user in our records"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response.ID = request.ID
	response.Balance = globals.Users[response.ID].Balance
	calculateStats(request.ID, &response)

	_ = json.NewEncoder(w).Encode(response)

}

func calculateStats(ID uint64, response *responses.GetUserResponse) {
	for _, deposit := range globals.UserDeposits[ID] {
		response.DepositCount++
		response.DepositSum += deposit.Amount
	}

	for _, transaction := range globals.UserTransactions[ID] {
		if transaction.Type == "Bet" {
			response.BetCount++
			response.BetSum += transaction.Amount
		}
		if transaction.Type == "Win" {
			response.WinCount++
			response.WinSum += transaction.Amount
		}
	}
}