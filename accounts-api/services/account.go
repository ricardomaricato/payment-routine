package services

import (
	"context"

	"github.com/ricardomaricato/payment-routine/accounts-api/models"
	"github.com/ricardomaricato/payment-routine/accounts-api/repositories"
)

// AccountService interface
type AccountService interface {
	CreateAccountService(ctx context.Context, account models.Account) (uint64, error)
}

// AccountServiceImpl implements AccountService
type AccountServiceImpl struct {
	accountRepository repositories.AccountRepository
}

// NewAccountService constructor
func NewAccountService(accountRepository repositories.AccountRepository) AccountService {
	return &AccountServiceImpl{accountRepository: accountRepository}
}

// CreateAccountService creates a new account
func (s *AccountServiceImpl) CreateAccountService(ctx context.Context, account models.Account) (uint64, error) {
	return s.accountRepository.CreateAccountRepository(ctx, account)
}
