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

// type AccountHandler interface {
// 	CreateAccountHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) error
// }

type AccountHandler struct {
	accountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (h *AccountHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Info("[CreateAccountHandler ioutil.ReadAll] Error reading body data")
		return
	}

	var account models.Account
	if err = json.Unmarshal(requestBody, &account); err != nil {
		log.Info("[CreateAccountHandler json.Unmarsahal] Error unmarshalling account")
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
