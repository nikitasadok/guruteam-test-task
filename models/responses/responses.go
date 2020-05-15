package responses

type ErrorResponse struct {
	Error string `json:"error"`
}

type AddUserResponse struct {
	Error string `json:"error"`
}
w
type AddDepositResponse struct {
	Error string `json:"error"`
	Balance float32 `json:"balance,omitempty"`
}

type GetUserResponse struct {
	ID uint64 `json:"id"`
	Error        string  `json:"error"`
	Balance      float32 `json:"balance"`
	DepositCount uint64  `json:"depositCount"`
	DepositSum   float32 `json:"depositSum"`
	BetCount     uint64 `json:"betCount"`
	BetSum float32 `json:"betSum"`
	WinCount uint64 `json:"winCount"`
	WinSum float32 `json:"winSum"`
}

type TransactionResponse struct {
	Error string `json:"error"`
	Balance float32 `json:"balance,omitempty"`
}