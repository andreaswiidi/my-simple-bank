package repository

import (
	"github.com/andreaswiidi/my-simple-bank/models"
	"gorm.io/gorm"
)

type TransferHistory struct {
	db *gorm.DB
}

func NewTransferHistoryRepository(db *gorm.DB) TransferHistory {
	return TransferHistory{
		db: db,
	}
}

func (th *TransferHistory) CreateTransferHistory(transfer models.TransfersHistory) (*models.TransfersHistory, error) {
	err := th.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&transfer)
		if result.Error != nil {
			// If there's an error, rollback the transaction
			return result.Error
		}
		return nil
	})

	if err != nil {
		return &transfer, err
	}

	return &transfer, nil

	// result := th.db.Create(transfer)
	// return transfer, result.Error
}

func (th *TransferHistory) GetTransferHistoryByTransactioID(accountId int64) (*models.TransfersHistory, error) {
	var TransactionHistory models.TransfersHistory
	result := th.db.Where("account_bank_id = ? AND is_deleted = ?", accountId, false).First(&TransactionHistory)
	if result.Error != nil {
		return nil, result.Error
	}
	return &TransactionHistory, nil
}
