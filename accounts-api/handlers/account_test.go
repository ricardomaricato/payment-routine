package handlers

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ricardomaricato/payment-routine/accounts-api/models"
	mock "github.com/ricardomaricato/payment-routine/accounts-api/services/mock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateAccountHandler_ShouldReturnsAccount_WhenSuccessful(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock.NewMockAccountService(ctrl)
	accountHandler := NewAccountHandler(service)

	service.EXPECT().CreateAccountService(context.Background(), gomock.Any()).DoAndReturn(func(account models.Account) error {
		account.Account_ID = 1
		account.AvailableCreditLimit = 1500.45
		account.AvailableWithDrawalLimit = 15000.0
		return nil
	})

	req, _ := http.NewRequest("POST", "/v1/accounts", bytes.NewBufferString(`
		{
   			"available_credit_limit": 1500.45,
    		"available_with_drawal_limit": 15000.0
		}
	`))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(accountHandler.CreateAccountHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_CreateAccountHandler_ShouldReturnsStatusInternalServerError_WhenFailsToCreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock.NewMockAccountService(ctrl)
	accountHandler := NewAccountHandler(service)

	service.EXPECT().CreateAccountService(context.Background(), gomock.Any()).Return(errors.New("Error"))

	req, _ := http.NewRequest("POST", "/v1/accounts", bytes.NewBufferString(`
		{
   			"available_credit_limit": 1500.45,
    		"available_with_drawal_limit": 15000.0
		}
	`))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(accountHandler.CreateAccountHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
