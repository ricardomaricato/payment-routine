package repositories

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/ricardomaricato/payment-routine/accounts-api/models"
	"gorm.io/gorm"
)

// AccountRepository interface
type AccountRepository interface {
	CreateAccountRepository(ctx context.Context, account models.Account) (uint64, error)
}

// AccountRepositoryImpl implements AccountRepository
type AccountRepositoryImpl struct {
	DB *gorm.DB
}

// NewAccountRepository constructor
func NewAccountRepository(DB *gorm.DB) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{DB}
}

// CreateAccountRepository creates a new account in database
func (r *AccountRepositoryImpl) CreateAccountRepository(ctx context.Context, account models.Account) (uint64, error) {
	if result := r.DB.WithContext(ctx).Create(&account).Last(&account); result.Error != nil {
		log.Errorf("[CreateAccountRepository] Failed to create account on DB. Error: %s", result.Error)
		return 0, result.Error
	}
	return account.Account_ID, nil
}
