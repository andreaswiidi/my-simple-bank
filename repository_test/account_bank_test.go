package repository_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// // newUser, _, err := createRandomUser()
// var newUserforAccount *models.User

// newUserforAccouunt = createRandomUser()

func TestCreateAccount(t *testing.T) {
	require.NotEmpty(t, userForTest)

	newBankAccount, err := testRepo.ACCOUNTBANK.CreateAccountBank(userForTest.ID)
	require.NoError(t, err)

	require.Equal(t, userForTest.ID, newBankAccount.UserID)
	require.NotZero(t, newBankAccount.ID)
	require.NotZero(t, newBankAccount.CreatedAt)
}

func TestGetUserandAccount(t *testing.T) {
	require.NotEmpty(t, userForTest)

	userFromRepo, err := testRepo.USER.FindUserByUsername(userForTest.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userFromRepo)

	if userFromRepo.AccountBank != nil {
		// userJSON, err := json.Marshal(userFromRepo)
		// require.NoError(t, err)
		// t.Log(string(userJSON))
		require.Equal(t, userFromRepo.AccountBank.UserID, userFromRepo.ID)
	}

	require.Equal(t, userForTest.Username, userFromRepo.Username)
	require.Equal(t, userForTest.Email, userFromRepo.Email)
	require.Equal(t, userForTest.Password, userFromRepo.Password)
	require.WithinDuration(t, userForTest.CreatedAt, userFromRepo.CreatedAt, time.Second)
}
