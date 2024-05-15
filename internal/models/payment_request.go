package models

import "time"

type PaymentRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Customer    string  `json:"customer"`
}

type CreatePaymentRequestResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID            int       `json:"id"`
		Domain        string    `json:"domain"`
		Amount        float64   `json:"amount"`
		Currency      string    `json:"currency"`
		DueDate       time.Time `json:"due_date"`
		HasInvoice    bool      `json:"has_invoice"`
		InvoiceNumber int       `json:"invoice_number"`
		Description   string    `json:"description"`
		LineItems     []struct {
			Name   string  `json:"name"`
			Amount float64 `json:"amount"`
		} `json:"line_items"`
		Tax []struct {
			Name   string  `json:"name"`
			Amount float64 `json:"amount"`
		} `json:"tax"`
		RequestCode      string    `json:"request_code"`
		Status           string    `json:"status"`
		Paid             bool      `json:"paid"`
		Metadata         any       `json:"metadata"`
		Notifications    []any     `json:"notifications"`
		OfflineReference string    `json:"offline_reference"`
		Customer         int       `json:"customer"`
		CreatedAt        time.Time `json:"created_at"`
	} `json:"data"`
}
