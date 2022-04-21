package routes

import (
	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/controllers"
	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/utils"
	"github.com/gorilla/mux"
)


var RegisterBankAccountRoutes = func (router *mux.Router)  {
	router.HandleFunc("/bank-account", controllers.CreateBankAccount).Methods("POST")
	router.HandleFunc("/bank-account/{username}/deposit", controllers.CreateBankAccountTransaction(utils.TransactionTypeCredit)).Methods("POST")
	router.HandleFunc("/bank-account/{username}/withdraw", controllers.CreateBankAccountTransaction(utils.TransactionTypeDebit)).Methods("POST")
}