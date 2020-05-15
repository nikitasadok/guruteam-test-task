package requests

type Request interface {
	GetToken() string
}

type AddUserRequest struct {
	ID uint64 `json:"id"`
	Balance float32 `json:"balance"`
	Token string `json:"token"`
}

func (addUserRequest AddUserRequest) GetToken() string {
	return addUserRequest.Token
}

type GetUserRequest struct {
	ID uint64 `json:"id"`
	Token string `json:"token"`
}

func (getUserRequest GetUserRequest) GetToken() string {
	return getUserRequest.Token
}

type AddDepositRequest struct {
	UserID uint64 `json:"userId"`
	DepositID uint64 `json:"depositId"`
	Amount float32 `json:"amount"`
	Token string `json:"token"`
}

func (addDepositRequest AddDepositRequest) GetToken() string {
	return addDepositRequest.Token
}

type TransactionRequest struct {
	UserID uint64 `json:"userId"`
	TransactionID uint64 `json:"transactionId"`
	Type string `json:"type"`
	Amount float32 `json:"amount"`
	Token string `json:"token"`
}

func (transactionRequest TransactionRequest) GetToken() string {
	return transactionRequest.Token
}