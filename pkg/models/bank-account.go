package models

import (
	"errors"

	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/config"
	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/utils"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type BankAccount struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
	BankAccountTransactions []BankAccountTransaction
}

type BankAccountTransaction struct {
	gorm.Model
	Note string `json:"note"`
	Amount int32 `json:"amount"`
	Type string `json:"type"`
	BankAccountID uint
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&BankAccount{})
	db.AutoMigrate(&BankAccountTransaction{})
}

func (bc *BankAccount) BeforeCreate(tx *gorm.DB) (err error) {
	var accountDetails BankAccount
	db.Where("username=?", bc.Username).Find(&accountDetails)
	if accountDetails.ID > 0 {
		return errors.New("username already exists")
	}
	if !utils.IsValidEmail(bc.Email) {
		return errors.New("invalid email address")
	}
	return
}

func (b *BankAccount) CreateBankAccount() (*BankAccount, error) {
	db.NewRecord(b)
	result := db.Create(&b)
	return b, result.Error
}


func (bt *BankAccountTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	if bt.Amount <= 0 {
		return errors.New("amount must be  greater than zero")
	}
	
	return
}

func (bt *BankAccountTransaction) CreateBankAccountTransaction(transactionType string, username string) (*BankAccountTransaction, error) {
	bt.Type = transactionType
	var accountDetails BankAccount
	db.Where("username=?", username).Find(&accountDetails)
	if accountDetails.ID == 0 {
		return bt,  errors.New("invalid username")
	}
	bt.BankAccountID = accountDetails.ID
	db.NewRecord(bt)
	result := db.Create(&bt)
	return bt, result.Error
}
