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
	var errorResponse responses.ErrorResponse
	var deposit entities.Deposit

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

	if !IsValidToken(request.Token) {
		errorResponse.Error = "The token is invalid!"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if globals.Users[request.UserID] == nil {
		errorResponse.Error = "There is no such user in our records"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if !isDepositeIDUnique(request.DepositID) {
		errorResponse.Error = "The deposite with such ID already exists!"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	deposit.TransferTime = time.Now()

	deposit.DepositID = request.DepositID
	deposit.Amount = request.Amount

	deposit.BalanceBefore = globals.Users[request.UserID].Balance
	deposit.BalanceAfter = deposit.BalanceBefore + deposit.Amount

	globals.UserDeposits[request.UserID] = append(globals.UserDeposits[request.UserID], &deposit)
	globals.Users[request.UserID].Balance = deposit.BalanceAfter
	print(globals.Users[request.UserID].ID)
	globals.RecentlyChangedUsers = append(globals.RecentlyChangedUsers, globals.Users[request.UserID])

	response.Balance = deposit.BalanceAfter

	_ = json.NewEncoder(w).Encode(response)
}

func IsValidToken(userToken string) bool {
	return userToken == globals.ServerToken
}

func isDepositeIDUnique (depositID uint64) bool {
	for _, deposits := range globals.UserDeposits {
		for _, deposit := range deposits {
			if deposit.DepositID == depositID {
				return false
			}
		}
	}
	return true
}
