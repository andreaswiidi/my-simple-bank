package repository

import (
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

func (u *User) FindAllUsers() ([]models.User, error) {
	var users []models.User
	result := u.db.Where("is_deleted = ?", false).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u *User) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := u.db.Preload("AccountBank").Where("username = ? AND is_deleted = ?", username, false).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *User) CreateUser(user models.User) (*models.User, error) {
	result := u.db.Create(&user)
	return &user, result.Error
}

func (u *User) EditUser(updatedUser *models.User) (*models.User, error) {
	// Save changes
	result := u.db.Save(updatedUser)
	if result.Error != nil {
		return updatedUser, result.Error
	}
	return updatedUser, nil
}

// func (u *User) Delete(username string) error {
// 	var user models.User

// 	user.IsDeleted = true

// 	result := u.db.Save(&user)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// 	// helper.ErrorPanic(result.Error)
// }
