package entities

import "time"

type User struct {
	ID uint64 `json:"id"`
	Balance float32
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