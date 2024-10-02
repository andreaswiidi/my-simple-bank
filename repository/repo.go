package repository

import "gorm.io/gorm"

type Repository struct {
	USER               User
	ACCOUNTBANK        AccountBank
	TRANSACTIONHISTORY TransactionHistory
	TRANSFERHISTORY    TransferHistory
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		USER:               NewUserRepository(db),
		ACCOUNTBANK:        NewAccountBankRepository(db),
		TRANSACTIONHISTORY: NewTransactionHistoryRepository(db),
		TRANSFERHISTORY:    NewTransferHistoryRepository(db),
	}
}
