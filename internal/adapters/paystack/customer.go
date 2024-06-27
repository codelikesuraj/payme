package paystack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	models "payme/internal/domain"
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

func (adapter *CustomerAdapter) CreateCustomer(customer models.Customer) (models.CreateCustomerResp, error) {
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

	resp, err := adapter.client.Do(req)
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

func (adapter *CustomerAdapter) FetchCustomer(customer string) (models.FetchCustomerResp, error) {
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

	resp, err := adapter.client.Do(req)
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
