package models

// Account of a client
type Account struct {
	Account_ID               uint64  `json:"account_id"`
	AvailableCreditLimit     float64 `json:"available_credit_limit"`
	AvailableWithDrawalLimit float64 `json:"available_with_drawal_limit"`
}
