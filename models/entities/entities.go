package entities

import "time"

type User struct {
	ID uint64 `json:"id"`
	Balance float32
}

type Statistics struct {
	DepositCount uint64
	DepositSum   float32
	BetCount     uint64
	BetSum float32
	WinCount uint64
	WinSum float32
}

type Deposit struct {
	UserID uint64
	DepositID uint64 `gorm:"primary_key"`
	Amount float32
	BalanceBefore float32
	BalanceAfter float32
	TransferTime time.Time
}

type Transaction struct {
	UserID uint64
	TransactionID uint64 `gorm:"primary_key"`
	Type string
	Amount float32
	BalanceBefore float32
	BalanceAfter float32
	TransactionTime time.Time
}