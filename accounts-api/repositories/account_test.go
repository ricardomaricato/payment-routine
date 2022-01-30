package repositories

import (
	"context"
	"testing"

	"github.com/ricardomaricato/payment-routine/accounts-api/config"
	"github.com/ricardomaricato/payment-routine/accounts-api/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_CreateAccountRepository_ShouldReturnsAccountID_WhenSuccessful(t *testing.T) {
	db, _ := gorm.Open(mysql.Open(config.DataBaseConectionString), &gorm.Config{})
	db.Migrator().DropTable(&models.Account{})
	db.AutoMigrate(&models.Account{})

	repo := NewAccountRepository(db)

	account := models.Account{
		AvailableCreditLimit:     500.0,
		AvailableWithDrawalLimit: 1000.0,
	}

	account.Account_ID, _ = repo.CreateAccountRepository(context.Background(), account)
	assert.NotNil(t, account)
	assert.Equal(t, uint64(0), account.Account_ID)
}
