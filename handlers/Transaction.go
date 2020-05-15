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

func Transaction(w http.ResponseWriter, r *http.Request) {
	var request requests.TransactionRequest
	var transaction entities.Transaction
	var errorResponse responses.ErrorResponse
	var response responses.TransactionResponse

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

	if request.Type != "Win" && request.Type != "Bet" {
		errorResponse.Error = "Cannot resolve the type of your transaction. It must be either Bet or Win"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if globals.Users[request.UserID] == nil {
		errorResponse.Error = "There is no such user in our records"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if request.Type == "Bet" && globals.Users[request.UserID].Balance < request.Amount {
		errorResponse.Error = "You don't have enough funds to make the bet. Please add a deposit to your account"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	transaction = entities.Transaction {
		TransactionID:   request.TransactionID,
		Type:            request.Type,
		Amount:          request.Amount,
		BalanceBefore:   globals.Users[request.UserID].Balance,
		TransactionTime: time.Now(),
	}

	if transaction.Type == "Win" {
		transaction.BalanceAfter = transaction.BalanceBefore + transaction.Amount
	} else {
		transaction.BalanceAfter = transaction.BalanceBefore - transaction.Amount
	}

	globals.UserTransactions[request.UserID] = append(globals.UserTransactions[request.UserID], &transaction)
	globals.Users[request.UserID].Balance = transaction.BalanceAfter

	response.Balance = transaction.BalanceAfter

	_ = json.NewEncoder(w).Encode(response)
}
