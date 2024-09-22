package repository_test

import (
	"testing"
	"time"

	"github.com/andreaswiidi/my-simple-bank/models"
	"github.com/andreaswiidi/my-simple-bank/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	// Arrange
	newUser := createRandomAccount()

	// Act: Create user
	testRepo.USER.CreateUser(newUser)

	// Assert: Verify the user exists
	var user models.User
	err := dbTest.Where("username = ?", newUser.Username).First(&user).Error
	assert.NoError(t, err)
	assert.Equal(t, newUser.Username, user.Username)
}

func createRandomAccount() models.User {
	randomName := util.RandomOwner()
	arg := models.User{
		Username:  randomName,
		Email:     randomName + "@mail.com",
		Password:  randomName,
		CreatedAt: time.Now(),
	}

	return arg
}
