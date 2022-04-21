package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/models"
	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/utils"
)


func CreateBankAccount(w http.ResponseWriter, r *http.Request) {
	NewBankAccount := &models.BankAccount{}
	utils.ParseBody(r, NewBankAccount)
	b := NewBankAccount.CreateBankAccount()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
