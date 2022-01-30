package services

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ricardomaricato/payment-routine/accounts-api/models"
	mock "github.com/ricardomaricato/payment-routine/accounts-api/repositories/mock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateAccountService_ShouldReturnsAccountID_WhenSuccessful(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()
	mockDao := mock.NewMockAccountRepository(crtl)
	service := NewAccountService(mockDao)

	account := models.Account{
		AvailableCreditLimit:     500.0,
		AvailableWithDrawalLimit: 1000.0,
	}

	mockDao.EXPECT().CreateAccountRepository(context.Background(), account).Return(nil, uint64(0)).AnyTimes()

	account.Account_ID, _ = service.CreateAccountService(context.Background(), account)
	assert.NotNil(t, account)
}
