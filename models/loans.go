package models

import "time"

// User Schema for users data storage
type Loans struct {
	Customer    string    `json:"customer"`
	Product     string    `json:"product"`
	Amount      string    `json:"amount"`
	Tenor       string    `json:"tenor"`
	Outstanding string    `json:"outstanding"`
	LoanStatus  string    `json:"loan_status"`
	InsertedAt  time.Time `json:"inserted_at"`
	LastUpdate  time.Time `json:"last_update"`
	//
}
