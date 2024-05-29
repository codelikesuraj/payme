package models

import "time"

type PaymentRequest struct {
	Amount           float64 `json:"amount"`
	Description      string  `json:"description"`
	Customer         string  `json:"customer"`
	SendNotification bool    `json:"send_notification"`
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

type ListPaymentRequestResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		ID            int       `json:"id"`
		Domain        string    `json:"domain"`
		Amount        float64   `json:"amount"`
		Currency      string    `json:"currency"`
		DueDate       time.Time `json:"due_date"`
		HasInvoice    bool      `json:"has_invoice"`
		InvoiceNumber int       `json:"invoice_number"`
		Description   string    `json:"description"`
		PdfURL        any       `json:"pdf_url"`
		LineItems     []struct {
			Name   string  `json:"name"`
			Amount float64 `json:"amount"`
		} `json:"line_items"`
		Tax []struct {
			Name   string  `json:"name"`
			Amount float64 `json:"amount"`
		} `json:"tax"`
		RequestCode      string `json:"request_code"`
		Status           string `json:"status"`
		Paid             bool   `json:"paid"`
		PaidAt           any    `json:"paid_at"`
		Metadata         any    `json:"metadata"`
		Notifications    []any  `json:"notifications"`
		OfflineReference string `json:"offline_reference"`
		Customer         struct {
			ID           int    `json:"id"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Email        string `json:"email"`
			CustomerCode string `json:"customer_code"`
			Phone        any    `json:"phone"`
			Metadata     struct {
				CallingCode string `json:"calling_code"`
			} `json:"metadata"`
			RiskAction               string `json:"risk_action"`
			InternationalFormatPhone any    `json:"international_format_phone"`
		} `json:"customer"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"data"`
	Meta struct {
		Total     int `json:"total"`
		Skipped   int `json:"skipped"`
		PerPage   any `json:"perPage"`
		Page      int `json:"page"`
		PageCount int `json:"pageCount"`
	} `json:"meta"`
}

type FetchPaymentRequestResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Transactions []any  `json:"transactions"`
		Domain       string `json:"domain"`
		RequestCode  string `json:"request_code"`
		Description  string `json:"description"`
		LineItems    []struct {
			Name   string `json:"name"`
			Amount int    `json:"amount"`
		} `json:"line_items"`
		Tax []struct {
			Name   string `json:"name"`
			Amount int    `json:"amount"`
		} `json:"tax"`
		Amount           int       `json:"amount"`
		Discount         any       `json:"discount"`
		Currency         string    `json:"currency"`
		DueDate          time.Time `json:"due_date"`
		Status           string    `json:"status"`
		Paid             bool      `json:"paid"`
		PaidAt           any       `json:"paid_at"`
		Metadata         any       `json:"metadata"`
		HasInvoice       bool      `json:"has_invoice"`
		InvoiceNumber    int       `json:"invoice_number"`
		OfflineReference string    `json:"offline_reference"`
		PdfURL           any       `json:"pdf_url"`
		Notifications    []any     `json:"notifications"`
		Archived         bool      `json:"archived"`
		Source           string    `json:"source"`
		PaymentMethod    any       `json:"payment_method"`
		Note             any       `json:"note"`
		AmountPaid       any       `json:"amount_paid"`
		ID               int       `json:"id"`
		Integration      int       `json:"integration"`
		Customer         struct {
			Transactions   []any  `json:"transactions"`
			Subscriptions  []any  `json:"subscriptions"`
			Authorizations []any  `json:"authorizations"`
			FirstName      string `json:"first_name"`
			LastName       string `json:"last_name"`
			Email          string `json:"email"`
			Phone          any    `json:"phone"`
			Metadata       struct {
				CallingCode string `json:"calling_code"`
			} `json:"metadata"`
			Domain       string    `json:"domain"`
			CustomerCode string    `json:"customer_code"`
			RiskAction   string    `json:"risk_action"`
			ID           int       `json:"id"`
			Integration  int       `json:"integration"`
			CreatedAt    time.Time `json:"createdAt"`
			UpdatedAt    time.Time `json:"updatedAt"`
		} `json:"customer"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
		PendingAmount int       `json:"pending_amount"`
	} `json:"data"`
}