package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testTaskGuru/globals"
	"testTaskGuru/models/entities"
	"testTaskGuru/models/requests"
	"testTaskGuru/models/responses"
	"time"
)

func AddDeposit (w http.ResponseWriter, r *http.Request) {
	var request requests.AddDepositRequest
	var response responses.AddDepositResponse
	var deposit entities.Deposit

	var err error

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Error = "Error reading your request into byte array"
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	err = json.Unmarshal(body, &request)

	if err != nil {
		response.Error = "Error unmarshalling the body of your request into a struct"
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	if !IsValidToken(request.Token) {
		response.Error = "The token is invalid!"
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	if globals.Users[request.UserID] == nil {
		response.Error = "There is no such user in our records"
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	deposit.TransferTime = time.Now()

	deposit.DepositID = request.DepositID
	deposit.Amount = request.Amount

	deposit.BalanceBefore = globals.Users[request.UserID].Balance
	deposit.BalanceAfter = deposit.BalanceBefore + deposit.Amount

	globals.UserDeposits[request.UserID] = append(globals.UserDeposits[request.UserID], &deposit)
	globals.Users[request.UserID].Balance = deposit.BalanceAfter

	response.Balance = deposit.BalanceAfter

	_ = json.NewEncoder(w).Encode(response)
}

func IsValidToken(userToken string) bool {
	return userToken == globals.ServerToken
}