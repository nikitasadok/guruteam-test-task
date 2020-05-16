package globals

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testTaskGuru/models/entities"
)

var Users map[uint64]*entities.User
var UserDeposits map[uint64][]*entities.Deposit
var UserTransactions map[uint64][]*entities.Transaction
var UserStatistics map[uint64][]*entities.Statistics

var RecentlyChangedUsers []*entities.User
var Db *gorm.DB

const ServerToken = "guru.team"

func InitApp() {
	Users = make(map[uint64]*entities.User)
	UserDeposits = make(map[uint64][]*entities.Deposit)
	UserTransactions = make(map[uint64][]*entities.Transaction)
	UserStatistics = make(map[uint64][]*entities.Statistics)
	Db, _ = gorm.Open("mysql", "root:nikita@/test_task?charset=utf8&parseTime=True&loc=Local")
	loadUserCacheFromDB()
	loadDepositCacheFromDB()
	loadTransactionCacheFromDB()
}

func UpdateDatabase() {
	for _,user := range RecentlyChangedUsers {
		print(user.ID)
		Db.Model(user).Where("id = ?", user.ID).Update("balance", user.Balance)
	}
}

func loadUserCacheFromDB() {
	var users []entities.User
	Db.Find(&users)
	for idx, user := range users {
		Users[user.ID] = &users[idx]
	}
}

func loadDepositCacheFromDB() {
	var deposits []entities.Deposit
	Db.Find(&deposits)
	for idx, deposit := range deposits {
		UserDeposits[deposit.UserID] = append(UserDeposits[deposit.UserID],&deposits[idx])
	}
}

func loadTransactionCacheFromDB() {
	var transactions []entities.Transaction
	Db.Find(&transactions)
	for idx, transaction := range transactions {
		UserTransactions[transaction.UserID] = append(UserTransactions[transaction.UserID], &transactions[idx])
	}
}

func calculateBalance() {
	for _, user := range Users {
		for  _, userDeposit := range UserDeposits[user.ID] {
			user.Balance += userDeposit.Amount
		}
		for _, userTransaction := range UserTransactions[user.ID] {
			if userTransaction.Type == "Win" {
				user.Balance += userTransaction.Amount
			}
			if userTransaction.Type == "Bet" {
				user.Balance -= userTransaction.Amount
			}
		}
	}
}





