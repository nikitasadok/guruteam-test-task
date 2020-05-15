package responses

type AddUserResponse struct {
	Error string `json:"error"`
}

type AddDepositResponse struct {
	Error string `json:"error"`
	Balance float32 `json:"balance,omitempty"`
}

type TransactionResponse struct {
	Error string `json:"error"`
	Balance float32 `json:"balance,omitempty"`
}