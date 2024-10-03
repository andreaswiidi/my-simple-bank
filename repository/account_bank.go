package repository

import (
	"github.com/andreaswiidi/my-simple-bank/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountBank struct {
	db *gorm.DB
}

func NewAccountBankRepository(db *gorm.DB) AccountBank {
	return AccountBank{
		db: db,
	}
}

func (ab *AccountBank) CreateAccountBank(acc models.AccountBank) (*models.AccountBank, error) {
	err := ab.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&acc)
		if result.Error != nil {
			// If there's an error, rollback the transaction
			return result.Error
		}
		return nil
	})
	if err != nil {
		return &acc, err
	}

	return &acc, nil
	// result := ab.db.Create(&acc)
	// return acc, result.Error
}

func (ab *AccountBank) UpdateAccountBank(updatedAccount *models.AccountBank) (*models.AccountBank, error) {
	// Save changes

	// Start the transaction
	err := ab.db.Transaction(func(tx *gorm.DB) error {
		var lockingAcc models.AccountBank
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", updatedAccount.ID).First(&lockingAcc).Error; err != nil {
			return err
		}

		result := tx.Save(updatedAccount)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})

	if err != nil {
		return updatedAccount, err
	}

	return updatedAccount, nil

	// result := u.db.Save(updatedAccount)
	// if result.Error != nil {
	// 	return updatedAccount, result.Error
	// }
	// return updatedAccount, nil
}

// func (ab *AccountBank) ReadAcc
