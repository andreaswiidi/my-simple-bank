package repository_test

import (
	"testing"
	"time"

	"github.com/andreaswiidi/my-simple-bank/models"
	"github.com/andreaswiidi/my-simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	newUser, arg, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	require.Equal(t, arg.Username, newUser.Username)
	require.Equal(t, arg.Email, newUser.Email)
	require.Equal(t, arg.Password, newUser.Password)

	require.NotZero(t, newUser.ID)
	require.NotZero(t, newUser.CreatedAt)
}

func createRandomUser() (*models.User, models.User, error) {
	randomName := util.RandomOwner()
	arg := models.User{
		FullName:  randomName,
		Username:  randomName,
		Email:     randomName + "@mail.com",
		Password:  randomName,
		CreatedAt: time.Now(),
	}

	// var user models.User
	user, err := testRepo.USER.CreateUser(arg)
	// err = dbTest.Where("username = ?", arg.Username).First(&user).Error

	return user, arg, err
}

func TestFindUserByUsername(t *testing.T) {
	newUser, _, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	userFromRepo, err := testRepo.USER.FindUserByUsername(newUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userFromRepo)

	require.Equal(t, newUser.Username, userFromRepo.Username)
	require.Equal(t, newUser.Email, userFromRepo.Email)
	require.Equal(t, newUser.Password, userFromRepo.Password)
	require.WithinDuration(t, newUser.CreatedAt, userFromRepo.CreatedAt, time.Second)
}

func TestUpdateUserByUsername(t *testing.T) {
	newUser, _, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	randomName := util.RandomOwner()
	newUser.FullName = randomName

	timeEdit := time.Now()
	newUser, err = testRepo.USER.EditUser(newUser)
	require.NoError(t, err)

	require.Equal(t, newUser.FullName, randomName)
	require.NotZero(t, newUser.UpdatedAt)
	require.WithinDuration(t, timeEdit, *newUser.UpdatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	newUser, _, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	newUser.IsDeleted = true

	newUser, err = testRepo.USER.EditUser(newUser)
	require.NoError(t, err)
	require.Equal(t, newUser.IsDeleted, true)

	userFromRepo, err := testRepo.USER.FindUserByUsername(newUser.Username)
	require.Error(t, err)
	require.Empty(t, userFromRepo)
}
