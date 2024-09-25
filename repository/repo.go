package repository

import "gorm.io/gorm"

type Repository struct {
	USER        User
	ACCOUNTBANK AccountBank
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		USER:        NewUserRepository(db),
		ACCOUNTBANK: NewAccountBankRepository(db),
	}
}
