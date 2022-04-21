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
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&BankAccount{})
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
