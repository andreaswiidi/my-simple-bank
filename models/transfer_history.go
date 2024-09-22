package models

import "time"

type TransfersHistory struct {
	ID         int64     `json:"id" gorm:"primary_key;auto_increment;not null;"`
	FromUserID int64     `json:"from_user_id" gorm:"not null;"`
	ToUserID   int64     `json:"to_user_id" gorm:"not null;"`
	Amount     int64     `json:"amount" gorm:"not null;"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null;"`
	FromUser   User      `gorm:"foreignKey:FromUserID"`
	ToUser     User      `gorm:"foreignKey:ToUserID"`
}
