package controllers

import (
	"net/http"

	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/models"
	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/utils"
	"github.com/gorilla/mux"
)


func CreateBankAccount(w http.ResponseWriter, r *http.Request) {
	NewBankAccount := &models.BankAccount{}
	utils.ParseBody(r, NewBankAccount)
	b, err := NewBankAccount.CreateBankAccount()
	utils.JSONReponse(w, b, err)
}

func CreateBankAccountTransactionCredit(w http.ResponseWriter, r *http.Request) {
	NewBankAccountTransaction := &models.BankAccountTransaction{}
	utils.ParseBody(r, NewBankAccountTransaction)
	params :=  mux.Vars(r)
	b, err := NewBankAccountTransaction.CreateBankAccountTransaction(utils.TransactionTypeCredit, params["username"])
	utils.JSONReponse(w, b, err)
}
