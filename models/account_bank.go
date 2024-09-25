package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountBank struct {
	gorm.Model
	ID                 int64                `json:"id" gorm:"primary_key;auto_increment;not null"`
	UserID             int64                `gorm:"not null;unique"`
	CreatedAt          time.Time            `json:"created_at" gorm:"not null;"`
	UpdatedAt          *time.Time           `json:"updated_at"`
	IsDeleted          bool                 `gorm:"not null;default:false"`
	TransactionHistory []TransactionHistory `gorm:"foreignKey:AccountBankID"`
}
