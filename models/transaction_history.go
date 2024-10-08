package models

import (
	"time"
)

type TransactionHistory struct {
	ID                int64            `json:"id" gorm:"primary_key;auto_increment;not null"`
	AccountBankID     int64            `json:"account_bank_id" gorm:"not null;"`
	Amount            int64            `json:"amount" gorm:"not null;"`
	TransactionType   string           `json:"transaction_type" gorm:"not null;"`
	TransferHistoryID int64            `gorm:"not null;" json:"transfer_history_id"`
	TransferHistory   TransfersHistory `gorm:"foreignKey:TransferHistoryID"`
	CreatedAt         time.Time        `json:"created_at" gorm:"not null;"`
}
