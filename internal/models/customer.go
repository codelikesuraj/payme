package models

import "time"

type Customer struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type CreateCustomerResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Email           string    `json:"email"`
		Integration     int       `json:"integration"`
		Domain          string    `json:"domain"`
		CustomerCode    string    `json:"customer_code"`
		ID              int       `json:"id"`
		Identified      bool      `json:"identified"`
		Identifications any       `json:"identifications"`
		CreatedAt       time.Time `json:"createdAt"`
		UpdatedAt       time.Time `json:"updatedAt"`
	} `json:"data"`
}

type FetchCustomerResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Transactions   []any `json:"transactions"`
		Subscriptions  []any `json:"subscriptions"`
		Authorizations []struct {
			AuthorizationCode string `json:"authorization_code"`
			Bin               string `json:"bin"`
			Last4             string `json:"last4"`
			ExpMonth          string `json:"exp_month"`
			ExpYear           string `json:"exp_year"`
			Channel           string `json:"channel"`
			CardType          string `json:"card_type"`
			Bank              string `json:"bank"`
			CountryCode       string `json:"country_code"`
			Brand             string `json:"brand"`
			Reusable          bool   `json:"reusable"`
			Signature         string `json:"signature"`
			AccountName       any    `json:"account_name"`
		} `json:"authorizations"`
		FirstName             any       `json:"first_name"`
		LastName              any       `json:"last_name"`
		Email                 string    `json:"email"`
		Phone                 any       `json:"phone"`
		Metadata              any       `json:"metadata"`
		Domain                string    `json:"domain"`
		CustomerCode          string    `json:"customer_code"`
		RiskAction            string    `json:"risk_action"`
		ID                    int       `json:"id"`
		Integration           int       `json:"integration"`
		CreatedAt             time.Time `json:"createdAt"`
		UpdatedAt             time.Time `json:"updatedAt"`
		CreatedAt0            time.Time `json:"created_at"`
		UpdatedAt0            time.Time `json:"updated_at"`
		TotalTransactions     int       `json:"total_transactions"`
		TotalTransactionValue []any     `json:"total_transaction_value"`
		DedicatedAccount      any       `json:"dedicated_account"`
		Identified            bool      `json:"identified"`
		Identifications       any       `json:"identifications"`
	} `json:"data"`
}
