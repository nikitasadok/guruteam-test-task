package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"testTaskGuru/handlers"
)

type Route struct {
	 Name string
	 Method string
	 Pattern string
	 HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	addUserRoute := Route{Name:"AddUser", Method:"POST", Pattern:"/user/create", HandlerFunc: handlers.AddUser }
	addDepositRoute := Route{Name:"AddDeposit", Method:"POST", Pattern:"/user/deposit", HandlerFunc: handlers.AddDeposit}
	transactionRoute := Route{Name:"Transaction", Method:"POST", Pattern:"/transaction", HandlerFunc: handlers.Transaction}

	router.Methods(addUserRoute.Method).Path(addUserRoute.Pattern).Name(addUserRoute.Name).
		HandlerFunc(addUserRoute.HandlerFunc)
	router.Methods(addDepositRoute.Method).Path(addDepositRoute.Pattern).Name(addDepositRoute.Name).
		HandlerFunc(addDepositRoute.HandlerFunc)
	router.Methods(transactionRoute.Method).Path(transactionRoute.Pattern).Name(transactionRoute.Name).
		HandlerFunc(transactionRoute.HandlerFunc)

	return router
}