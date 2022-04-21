package models

import (
	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type BankAccount struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&BankAccount{})
}

func (b *BankAccount) CreateBankAccount() *BankAccount {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
