package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testTaskGuru/commons"
	"testTaskGuru/globals"
	"testTaskGuru/models/entities"
	"testTaskGuru/models/requests"
	"testTaskGuru/models/responses"
)

func GetUser (w http.ResponseWriter, r *http.Request) {
	var request requests.GetUserRequest
	var response responses.GetUserResponse
	var errorResponse responses.ErrorResponse
	var stats entities.Statistics

	var err error

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		errorResponse.Error = "Error reading your request into byte array"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = json.Unmarshal(body, &request)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		errorResponse.Error = "Error unmarshalling the body of your request into a struct"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if !commons.IsValidToken(request.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		errorResponse.Error = "The token is invalid!"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if globals.Users[request.ID] == nil {
		w.WriteHeader(http.StatusConflict)
		errorResponse.Error = "There is no such user in our records"
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response.ID = request.ID
	response.Balance = globals.Users[response.ID].Balance
	calculateStats(request.ID, &stats)
	globals.UserStatistics[request.ID] = append(globals.UserStatistics[request.ID], &stats)

	response.WinCount = stats.WinCount
	response.WinSum = stats.WinSum
	response.BetCount = stats.BetCount
	response.BetSum = stats.BetSum
	response.DepositCount = stats.DepositCount
	response.DepositSum = stats.DepositSum

	_ = json.NewEncoder(w).Encode(response)

}

func calculateStats(ID uint64, stats *entities.Statistics) {
	for _, deposit := range globals.UserDeposits[ID] {
		stats.DepositCount++
		stats.DepositSum += deposit.Amount
	}

	for _, transaction := range globals.UserTransactions[ID] {
		if transaction.Type == "Bet" {
			stats.BetCount++
			stats.BetSum += transaction.Amount
		}
		if transaction.Type == "Win" {
			stats.WinCount++
			stats.WinSum += transaction.Amount
		}
	}
}