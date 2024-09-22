package repository_test

import (
	"log"
	"os"
	"testing"

	"github.com/andreaswiidi/my-simple-bank/config"
	"github.com/andreaswiidi/my-simple-bank/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var testRepo repository.Repository
var dbTest *gorm.DB

func TestMain(m *testing.M) {
	dbTest = config.ConnectDataBase()
	log.Printf("Database connection: %v", dbTest)

	if dbTest == nil {
		panic("Database connection is nil")
	}

	testRepo = repository.NewRepository(dbTest)

	os.Exit(m.Run())
}

func TestOpenConnection(t *testing.T) {
	// dbTest := config.ConnectDataBase()
	t.Logf("Database connection: %v", dbTest)
	assert.NotNil(t, dbTest)
}
