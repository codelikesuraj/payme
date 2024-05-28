package adapters

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"payme/internal/models"
)

const (
	BASE_URL            = "https://api.paystack.co"
	CUSTOMER_URL        = "customer"
	PAYMENT_REQUEST_URL = "paymentrequest"
)

type PaystackAPIAdapter struct {
	httpClient *http.Client
}

func NewPaystackAPIAdapter() *PaystackAPIAdapter {
	return &PaystackAPIAdapter{
		httpClient: &http.Client{},
	}
}

func (adapter *PaystackAPIAdapter) CreateCustomer(customer models.Customer) (models.CreateCustomerResp, error) {
	jsonData, err := json.Marshal(customer)
	if err != nil {
		return models.CreateCustomerResp{}, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", BASE_URL, CUSTOMER_URL),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return models.CreateCustomerResp{}, err
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("PAYSTACK_SK"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := adapter.httpClient.Do(req)
	if err != nil {
		return models.CreateCustomerResp{}, err
	}
	defer resp.Body.Close()

	var body models.CreateCustomerResp
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return models.CreateCustomerResp{}, err
	}

	if !body.Status {
		return models.CreateCustomerResp{}, errors.New(body.Message)
	}

	return body, nil
}

func (adapter *PaystackAPIAdapter) FetchCustomer(customer string) (models.FetchCustomerResp, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s/%s", BASE_URL, CUSTOMER_URL, customer),
		nil,
	)
	if err != nil {
		return models.FetchCustomerResp{}, err
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("PAYSTACK_SK"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := adapter.httpClient.Do(req)
	if err != nil {
		return models.FetchCustomerResp{}, err
	}
	defer resp.Body.Close()

	var body models.FetchCustomerResp
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return models.FetchCustomerResp{}, err
	}

	if !body.Status {
		return models.FetchCustomerResp{}, errors.New(body.Message)
	}

	return body, nil
}

func (adapter *PaystackAPIAdapter) CreatePaymentRequest(pr models.PaymentRequest) (models.CreatePaymentRequestResp, error) {
	jsonData, err := json.Marshal(pr)
	if err != nil {
		return models.CreatePaymentRequestResp{}, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", BASE_URL, PAYMENT_REQUEST_URL),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return models.CreatePaymentRequestResp{}, err
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("PAYSTACK_SK"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := adapter.httpClient.Do(req)
	if err != nil {
		return models.CreatePaymentRequestResp{}, err
	}
	defer resp.Body.Close()

	var body models.CreatePaymentRequestResp
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return models.CreatePaymentRequestResp{}, err
	}

	if !body.Status {
		return models.CreatePaymentRequestResp{}, errors.New(body.Message)
	}

	return body, nil
}

func (adapter *PaystackAPIAdapter) ListPaymentRequest(flag models.ListFlag) (models.ListPaymentRequestResp, error) {
	url := fmt.Sprintf("%s/%s", BASE_URL, PAYMENT_REQUEST_URL)

	if flag.Last {
		url += "/?perPage=1"
	} else if flag.Count >= 1 {
		url += fmt.Sprintf("/?perPage=%d", flag.Count)
	}

	if flag.Page > 0 {
		url += fmt.Sprintf("&page=%d", flag.Page)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return models.ListPaymentRequestResp{}, err
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("PAYSTACK_SK"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := adapter.httpClient.Do(req)
	if err != nil {
		return models.ListPaymentRequestResp{}, err
	}
	defer resp.Body.Close()

	var body models.ListPaymentRequestResp

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return models.ListPaymentRequestResp{}, err
	}

	if !body.Status {
		return models.ListPaymentRequestResp{}, errors.New(body.Message)
	}

	return body, nil
}
