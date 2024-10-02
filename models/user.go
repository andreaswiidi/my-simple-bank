package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          int64        `json:"id" gorm:"primary_key;auto_increment;not null"`
	FullName    string       `gorm:"not null;" json:"fullname"`
	Username    string       `gorm:"not null;unique" json:"username"`
	Email       string       `gorm:"not null;unique" json:"email"`
	Password    string       `gorm:"not null;" json:"password"`
	CreatedAt   time.Time    `json:"created_at" gorm:"not null;"`
	UpdatedAt   *time.Time   `json:"updated_at"`
	IsDeleted   bool         `gorm:"not null;default:false"`
	AccountBank *AccountBank `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"accountBank"`
}
