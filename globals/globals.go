package globals

import "testTaskGuru/models/entities"

var Users map[uint64]*entities.User
var UserDeposits map[uint64][]*entities.Deposit
var UserTransactions map[uint64][]*entities.Transaction

const ServerToken = "guru.team"

func InitApp() {
	Users = make(map[uint64]*entities.User)
	UserDeposits = make(map[uint64][]*entities.Deposit)
	UserTransactions = make(map[uint64][]*entities.Transaction)
}


