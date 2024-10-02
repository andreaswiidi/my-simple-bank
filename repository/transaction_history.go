package repository

import (
	"github.com/andreaswiidi/my-simple-bank/models"
	"gorm.io/gorm"
)

type TransactionHistory struct {
	db *gorm.DB
}

func NewTransactionHistoryRepository(db *gorm.DB) TransactionHistory {
	return TransactionHistory{
		db: db,
	}
}

func (th *TransactionHistory) CreateTransaction(transaction *models.TransactionHistory) (*models.TransactionHistory, error) {
	result := th.db.Create(transaction)
	return transaction, result.Error
}

func (th *TransactionHistory) GetTransactionHistoryByAccountId(accountId int64) (*models.TransactionHistory, error) {
	var TransactionHistory models.TransactionHistory
	result := th.db.Where("account_bank_id = ? AND is_deleted = ?", accountId, false).First(&TransactionHistory)
	if result.Error != nil {
		return nil, result.Error
	}
	return &TransactionHistory, nil
}
