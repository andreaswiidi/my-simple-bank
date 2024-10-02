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

func (ab *AccountBank) CreateAccountBank(acc *models.AccountBank) (*models.AccountBank, error) {
	result := ab.db.Create(&acc)
	return acc, result.Error
}

func (u *AccountBank) UpdateAccountBank(updatedAccount *models.AccountBank) (*models.AccountBank, error) {
	// Save changes
	result := u.db.Save(updatedAccount)
	if result.Error != nil {
		return updatedAccount, result.Error
	}
	return updatedAccount, nil
}

// func (ab *AccountBank) ReadAcc
