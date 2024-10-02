package repository_test

import (
	"testing"
	"time"

	"github.com/andreaswiidi/my-simple-bank/models"
	"github.com/andreaswiidi/my-simple-bank/util"
	"github.com/stretchr/testify/require"
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

			argAcc := &models.AccountBank{
				UserID:   userForTransactionA.ID,
				Balance:  randomNumber,
				Currency: randomCurrency,
			}

			argAcc, err = testRepo.ACCOUNTBANK.CreateAccountBank(argAcc)
			if err != nil {
				return err
			}
			accountForTransactionA = argAcc

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

			argAccB := &models.AccountBank{
				UserID:   userForTransactionB.ID,
				Balance:  randomNumber,
				Currency: randomCurrency,
			}

			argAccB, err = testRepo.ACCOUNTBANK.CreateAccountBank(argAccB)
			if err != nil {
				return err
			}
			accountForTransactionB = argAccB

		}
	}

	return err
}

func TestTransaction(t *testing.T) {
	err := prepareTransaction()
	require.NoError(t, err)
	amount := int64(10)

	//make transfer
	makeTransfer := &models.TransfersHistory{
		FromAccountID: accountForTransactionB.UserID,
		ToAccountID:   accountForTransactionA.UserID,
		Amount:        amount,
	}

	makeTransfer, err = testRepo.TRANSFERHISTORY.CreateTransferHistory(makeTransfer)
	require.NoError(t, err)
	require.NotEmpty(t, makeTransfer)

	makeTransactionA := &models.TransactionHistory{
		AccountBankID:     accountForTransactionA.ID,
		TransactionType:   util.TRANSACTION_TYPE_TRANSFER,
		TransferHistoryID: makeTransfer.ID,
		Amount:            amount,
	}
	makeTransactionA, err = testRepo.TRANSACTIONHISTORY.CreateTransaction(makeTransactionA)
	require.NoError(t, err)
	require.NotEmpty(t, makeTransactionA)

	accountForTransactionA.Balance = accountForTransactionA.Balance + amount
	accountForTransactionA, err = testRepo.ACCOUNTBANK.UpdateAccountBank(accountForTransactionA)
	require.NoError(t, err)

	makeTransactionB := &models.TransactionHistory{
		AccountBankID:     accountForTransactionA.ID,
		TransactionType:   util.TRANSACTION_TYPE_TRANSFER,
		TransferHistoryID: makeTransfer.ID,
		Amount:            -amount,
	}
	makeTransactionB, err = testRepo.TRANSACTIONHISTORY.CreateTransaction(makeTransactionB)
	require.NoError(t, err)
	require.NotEmpty(t, makeTransactionB)

	accountForTransactionA, err = testRepo.ACCOUNTBANK.UpdateAccountBank(accountForTransactionA)
	require.NoError(t, err)

	// errs := make(chan error)
	// results := make(chan models.TransfersHistory)
	// go func() {
	// 	ctx := context.WithValue(context.Background(), txKey, txName)
	// 	result, err := store.TransferTx(ctx, sqlc.TransferTxParams{
	// 		FromAccountID: account1.ID,
	// 		ToAccountID:   account2.ID,
	// 		Amount:        amount,
	// 	})

	// 	errs <- err
	// 	results <- result
	// }()
}
