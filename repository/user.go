package repository

import (
	"github.com/andreaswiidi/my-simple-bank/helper"
	"github.com/andreaswiidi/my-simple-bank/models"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return User{
		db: db,
	}
}

func (u *User) CreateUser(user models.User) {
	result := u.db.Create(&user)
	helper.ErrorPanic(result.Error)
}

func (t *User) Delete(username string) {
	var tags models.User
	result := t.db.Where("username = ?", username).Delete(&tags)
	helper.ErrorPanic(result.Error)
}
