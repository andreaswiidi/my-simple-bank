package repository

import "gorm.io/gorm"

type Repository struct {
	USER User
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		USER: NewUserRepository(db),
	}
}
