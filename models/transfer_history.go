package models

import "time"

type TransfersHistory struct {
	ID            int64     `json:"id" gorm:"primary_key;auto_increment;not null;"`
	FromAccountID int64     `json:"from_account_id" gorm:"not null;"`
	ToAccountID   int64     `json:"to_account_id" gorm:"not null;"`
	Amount        int64     `json:"amount" gorm:"not null;"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null;"`
}
