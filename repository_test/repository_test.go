package repository_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/andreaswiidi/my-simple-bank/config"
	"github.com/andreaswiidi/my-simple-bank/models"
	"github.com/andreaswiidi/my-simple-bank/repository"
	"github.com/andreaswiidi/my-simple-bank/util"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var testRepo repository.Repository
var dbTest *gorm.DB
var userForTest *models.User

func TestMain(m *testing.M) {
	var err error
	dbTest = config.ConnectDataBase()
	log.Printf("Database connection: %v", dbTest)

	if dbTest == nil {
		panic("Database connection is nil")
	}

	testRepo = repository.NewRepository(dbTest)

	randomName := util.RandomOwner()

	userForTest = &models.User{
		FullName:  randomName,
		Username:  randomName,
		Email:     randomName + "@mail.com",
		Password:  randomName,
		CreatedAt: time.Now(),
	}
	createdUser, err := testRepo.USER.CreateUser(userForTest)

	userForTest = createdUser

	if err != nil {
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestOpenConnection(t *testing.T) {
	// dbTest := config.ConnectDataBase()
	// t.Logf("Database connection: %v", dbTest)
	assert.NotNil(t, dbTest)
}
