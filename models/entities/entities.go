package entities

import "time"

type User struct {
	ID uint64 `json:"id"`
	Balance float32
}

type Statistics struct {
	DepositCount uint64  `json:"depositCount"`
	DepositSum   float32 `json:"depositSum"`
	BetCount     uint64 `json:"betCount"`
	BetSum float32 `json:"betSum"`
	WinCount uint64 `json:"winCount"`
	WinSum float32 `json:"winSum"`
}

type Deposit struct {
	DepositID uint64
	Amount float32
	BalanceBefore float32
	BalanceAfter float32
	TransferTime time.Time
}

type Transaction struct {
	TransactionID uint64
	Type string
	Amount float32
	BalanceBefore float32
	BalanceAfter float32
	TransactionTime time.Time
}