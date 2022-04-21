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
	Name string `valid:"required" json:"name"`
	Email string `valid:"required" json:"email"`
	Username string `valid:"required" json:"username"`
	BankAccountTransactions []BankAccountTransaction
}

type BankAccountTransaction struct {
	gorm.Model
	Note string `json:"note"`
	Amount int32 `valid:"required" json:"amount"`
	Type string `valid:"required" json:"type"`
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

func (bt *BankAccountTransaction) CreateBankAccountTransaction(transactionType string, username string) (*BankAccountTransaction, error) {
	bt.Type = transactionType
	var accountDetails BankAccount
	db.Where("username=?", username).Find(&accountDetails)
	if accountDetails.ID == 0 {
		return bt,  errors.New("invalid username")
	}

	bt.BankAccountID = accountDetails.ID
	if bt.Amount <= 0 {
		return bt, errors.New("amount must be greater than zero")
	}

	if bt.Type == utils.TransactionTypeDebit {
		balance := GetBankAccountBalance(accountDetails.Username)
		if balance < bt.Amount {
			return bt, errors.New("balance insufficient")
		}
		bt.Amount = bt.Amount * -1
	}

	db.NewRecord(bt)
	result := db.Create(&bt)
	return bt, result.Error
}

func  GetBankAccountBalance(username string) (int32) {
	var balance int32
	query := db.Model(&BankAccountTransaction{}).Select("sum(amount) as balance")
	query = query.Joins("left join bank_accounts on bank_accounts.id = bank_account_transactions.bank_account_id")
	query.Where("bank_accounts.username = ?", username).Row().Scan(&balance)
	return balance
}