package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountBank struct {
	gorm.Model
	ID                 int64                 `json:"id" gorm:"primary_key;auto_increment;not null"`
	UserID             int64                 `gorm:"not null" json:"user_id"`
	Balance            int64                 `gorm:"not null;default:0" json:"balance"`
	Currency           string                `gorm:"not null" json:"currency"`
	CreatedAt          time.Time             `json:"created_at" gorm:"not null;"`
	UpdatedAt          *time.Time            `json:"updated_at"`
	IsDeleted          bool                  `gorm:"not null;default:false" json:"is_deleted"`
	TransactionHistory *[]TransactionHistory `gorm:"foreignKey:AccountBankID"`
}
