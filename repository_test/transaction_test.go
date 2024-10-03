package repository_test

import (
	// "fmt"
	"testing"
	"time"

	"github.com/andreaswiidi/my-simple-bank/models"
	"github.com/andreaswiidi/my-simple-bank/util"
	"github.com/stretchr/testify/require"
	// "gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

var userForTransactionA *models.User
var userForTransactionB *models.User

var accountForTransactionA *models.AccountBank
var accountForTransactionB *models.AccountBank

func prepareTransaction() error {
	var err error
	if userForTransactionA == nil {
		randomName := util.RandomOwner()

		arg := &models.User{
			FullName:  randomName,
			Username:  randomName,
			Email:     randomName + "@mail.com",
			Password:  randomName,
			CreatedAt: time.Now(),
		}
		arg, err = testRepo.USER.CreateUser(arg)
		if err != nil {
			return err
		}
		userForTransactionA = arg

		if accountForTransactionA == nil {
			randomNumber := util.RandomMoney()
			randomCurrency := util.RandomCurrency()

			argAcc := models.AccountBank{
				UserID:   userForTransactionA.ID,
				Balance:  randomNumber,
				Currency: randomCurrency,
			}

			accountForTransactionA, err = testRepo.ACCOUNTBANK.CreateAccountBank(argAcc)
			if err != nil {
				return err
			}

		}

	}

	if userForTransactionB == nil {
		randomName := util.RandomOwner()

		argB := &models.User{
			FullName:  randomName,
			Username:  randomName,
			Email:     randomName + "@mail.com",
			Password:  randomName,
			CreatedAt: time.Now(),
		}
		argB, err = testRepo.USER.CreateUser(argB)
		if err != nil {
			return err
		}
		userForTransactionB = argB

		if accountForTransactionB == nil {
			randomNumber := util.RandomMoney()
			randomCurrency := util.RandomCurrency()

			argAccB := models.AccountBank{
				UserID:   userForTransactionB.ID,
				Balance:  randomNumber,
				Currency: randomCurrency,
			}

			accountForTransactionB, err = testRepo.ACCOUNTBANK.CreateAccountBank(argAccB)
			if err != nil {
				return err
			}
		}
	}

	return err
}

func TestTransaction(t *testing.T) {
	err := prepareTransaction()
	require.NoError(t, err)
	amount := int64(10)

	t.Logf("Money Account A : %d", accountForTransactionA.Balance)
	t.Logf("Money Account B : %d", accountForTransactionB.Balance)

	//make transfer
	makeTransfer := models.TransfersHistory{
		FromAccountID: accountForTransactionB.UserID,
		ToAccountID:   accountForTransactionA.UserID,
		Amount:        amount,
	}

	transferHistory, err := testRepo.TRANSFERHISTORY.CreateTransferHistory(makeTransfer)
	require.NoError(t, err)
	require.NotEmpty(t, transferHistory)

	makeTransactionA := models.TransactionHistory{
		AccountBankID:     accountForTransactionA.ID,
		TransactionType:   util.TRANSACTION_TYPE_TRANSFER,
		TransferHistoryID: transferHistory.ID,
		Amount:            amount,
	}
	transactionA, err := testRepo.TRANSACTIONHISTORY.CreateTransaction(makeTransactionA)
	require.NoError(t, err)
	require.NotEmpty(t, transactionA)

	accountForTransactionA.Balance = accountForTransactionA.Balance + amount
	accountForTransactionA, err = testRepo.ACCOUNTBANK.UpdateAccountBank(accountForTransactionA)
	// t.Logf("Money Account A Update : %d", accountForTransactionA.Balance)
	require.NoError(t, err)

	makeTransactionB := models.TransactionHistory{
		AccountBankID:     accountForTransactionA.ID,
		TransactionType:   util.TRANSACTION_TYPE_TRANSFER,
		TransferHistoryID: transferHistory.ID,
		Amount:            -amount,
	}
	transactionB, err := testRepo.TRANSACTIONHISTORY.CreateTransaction(makeTransactionB)
	require.NoError(t, err)
	require.NotEmpty(t, transactionB)

	accountForTransactionB.Balance = accountForTransactionB.Balance - amount
	accountForTransactionB, err := testRepo.ACCOUNTBANK.UpdateAccountBank(accountForTransactionB)
	// t.Logf("Money Account B Update : %d", accountForTransactionB.Balance)
	require.NoError(t, err)

	t.Logf("Money Account A After : %d", accountForTransactionA.Balance)
	t.Logf("Money Account B After : %d", accountForTransactionB.Balance)
}

func TestRaceConditionTransaction(t *testing.T) {
	err := prepareTransaction()
	require.NoError(t, err)
	amount := int64(10)
	count := 5

	errs := make(chan error)
	results := make(chan models.TransfersHistory)

	for i := 0; i < count; i++ {
		go func() {

			accountForTransactionA.Balance = accountForTransactionA.Balance + amount
			accountForTransactionA, err = testRepo.ACCOUNTBANK.UpdateAccountBank(accountForTransactionA)
			// t.Logf("Money Account A Update : %d", accountForTransactionA.Balance)
			require.NoError(t, err)

			accountForTransactionB.Balance = accountForTransactionB.Balance - amount
			accountForTransactionB, err := testRepo.ACCOUNTBANK.UpdateAccountBank(accountForTransactionB)
			// t.Logf("Money Account B Update : %d", accountForTransactionB.Balance)
			require.NoError(t, err)

			//make transfer
			makeTransfer := models.TransfersHistory{
				FromAccountID: accountForTransactionB.UserID,
				ToAccountID:   accountForTransactionA.UserID,
				Amount:        amount,
			}

			transferHistory, err := testRepo.TRANSFERHISTORY.CreateTransferHistory(makeTransfer)
			require.NoError(t, err)
			require.NotEmpty(t, transferHistory)

			makeTransactionA := models.TransactionHistory{
				AccountBankID:     accountForTransactionA.ID,
				TransactionType:   util.TRANSACTION_TYPE_TRANSFER,
				TransferHistoryID: transferHistory.ID,
				Amount:            amount,
			}
			transactionA, err := testRepo.TRANSACTIONHISTORY.CreateTransaction(makeTransactionA)
			require.NoError(t, err)
			require.NotEmpty(t, transactionA)

			makeTransactionB := models.TransactionHistory{
				AccountBankID:     accountForTransactionA.ID,
				TransactionType:   util.TRANSACTION_TYPE_TRANSFER,
				TransferHistoryID: transferHistory.ID,
				Amount:            -amount,
			}
			transactionB, err := testRepo.TRANSACTIONHISTORY.CreateTransaction(makeTransactionB)
			require.NoError(t, err)
			require.NotEmpty(t, transactionB)

			errs <- err
			results <- *transferHistory
		}()
	}

	for i := 0; i < count; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
		t.Logf("Money Account A After : %d", accountForTransactionA.Balance)
		t.Logf("Money Account B After : %d", accountForTransactionB.Balance)
	}

}

// func makeTransaction(toAccount *models.AccountBank, fromAccount *models.AccountBank,ammout int64) error {
// 	//make transfer
// 	return dbTest.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", fromAccount.ID).First(&fromAccount).Error; err != nil {
// 			return fmt.Errorf("failed to find and lock from account: %w", err)
// 		}

// 		if fromAccount.Balance < ammout {
//             return fmt.Errorf("insufficient balance in from account")
//         }
// 		// Lock the toAccount for update
//         if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", toAccount.ID).First(&toAccount).Error; err != nil {
//             return fmt.Errorf("failed to find and lock to account: %w", err)
//         }

// 		fromAccount.Balance -= ammout
// 		if err := tx.Save(&fromAccount).Error; err != nil {
//             return fmt.Errorf("failed to deduct amount from fromAccount: %w", err)
//         }
// 	})
// }
