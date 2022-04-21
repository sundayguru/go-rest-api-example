package controllers

import (
	"net/http"

	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/models"
	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/utils"
)


func CreateBankAccount(w http.ResponseWriter, r *http.Request) {
	NewBankAccount := &models.BankAccount{}
	utils.ParseBody(r, NewBankAccount)
	b, err := NewBankAccount.CreateBankAccount()
	utils.JSONReponse(w, b, err)
}
