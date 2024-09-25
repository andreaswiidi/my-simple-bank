package repository

import (
	"github.com/andreaswiidi/my-simple-bank/models"
	"gorm.io/gorm"
)

type AccountBank struct {
	db *gorm.DB
}

func NewAccountBankRepository(db *gorm.DB) AccountBank {
	return AccountBank{
		db: db,
	}
}

func (ab *AccountBank) CreateAccountBank(userID int64) (*models.AccountBank, error) {
	var account models.AccountBank
	account.UserID = userID
	result := ab.db.Create(&account)
	return &account, result.Error
}
