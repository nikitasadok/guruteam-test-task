package requests

type AddUserRequest struct {
	ID uint64 `json:"id"`
	Balance float32 `json:"balance"`
	Token string `json:"token"`
}

type AddDepositRequest struct {
	UserID uint64 `json:"userId"`
	DepositID uint64 `json:"depositId"`
	Amount float32 `json:"amount"`
	Token string `json:"token"`
}

type TransactionRequest struct {
	UserID uint64 `json:"userId"`
	TransactionID uint64 `json:"transactionId"`
	Type string `json:"type"`
	Amount float32 `json:"amount"`
	Token string `json:"token"`
}