package routes

import (
	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/controllers"
	"github.com/gorilla/mux"
)


var RegisterBankAccountRoutes = func (router *mux.Router)  {
	router.HandleFunc("/bank-account", controllers.CreateBankAccount).Methods("POST")
}