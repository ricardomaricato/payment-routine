package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/ricardomaricato/payment-routine/accounts-api/models"
	"github.com/ricardomaricato/payment-routine/accounts-api/responses"
	"github.com/ricardomaricato/payment-routine/accounts-api/services"
)

// // AccountHandler interface
type AccountHandler interface {
	CreateAccountHandler(w http.ResponseWriter, r *http.Request)
}

// AccountHandler implements AccountHandler
type AccountHandlerImpl struct {
	accountService services.AccountService
}

// NewAccountHandler returns constructor
func NewAccountHandler(accountService services.AccountService) AccountHandler {
	return &AccountHandlerImpl{accountService: accountService}
}

// CreateAccountHandler creates a new account
func (h *AccountHandlerImpl) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Info("[CreateAccountHandler ReadAll] Error reading body data")
		responses.Err(w, http.StatusBadGateway, err)
		return
	}

	var account models.Account
	if err = json.Unmarshal(requestBody, &account); err != nil {
		log.Info("[CreateAccountHandler Unmarsahal] Error unmarshalling account")
		responses.Err(w, http.StatusBadGateway, err)
		return
	}

	account.Account_ID, err = h.accountService.CreateAccountService(r.Context(), account)
	if err != nil {
		log.Info("[CreateAccountHandler CreateAccountService] Error creating account")
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, account)
}
