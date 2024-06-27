package paystack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"payme/internal/domain"
)

type CustomerAdapter struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func NewCustomerAdapter(apiKey string) *CustomerAdapter {
	return &CustomerAdapter{
		apiKey:  apiKey,
		baseURL: BASE_URL,
		client:  &http.Client{},
	}
}

func (adapter *CustomerAdapter) CreateCustomer(customer domain.Customer) (domain.CreateCustomerResp, error) {
	jsonData, err := json.Marshal(customer)
	if err != nil {
		return domain.CreateCustomerResp{}, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", BASE_URL, CUSTOMER_URL),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return domain.CreateCustomerResp{}, err
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("PAYSTACK_SK"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := adapter.client.Do(req)
	if err != nil {
		return domain.CreateCustomerResp{}, err
	}
	defer resp.Body.Close()

	var body domain.CreateCustomerResp
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return domain.CreateCustomerResp{}, err
	}

	if !body.Status {
		return domain.CreateCustomerResp{}, errors.New(body.Message)
	}

	return body, nil
}

func (adapter *CustomerAdapter) FetchCustomer(customer string) (domain.FetchCustomerResp, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s/%s", BASE_URL, CUSTOMER_URL, customer),
		nil,
	)
	if err != nil {
		return domain.FetchCustomerResp{}, err
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("PAYSTACK_SK"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := adapter.client.Do(req)
	if err != nil {
		return domain.FetchCustomerResp{}, err
	}
	defer resp.Body.Close()

	var body domain.FetchCustomerResp
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return domain.FetchCustomerResp{}, err
	}

	if !body.Status {
		return domain.FetchCustomerResp{}, errors.New(body.Message)
	}

	return body, nil
}
