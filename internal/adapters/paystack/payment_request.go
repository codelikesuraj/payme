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

type PaymentRequestAdapter struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func NewPaymentRequestAdapter(apiKey string) *PaymentRequestAdapter {
	return &PaymentRequestAdapter{
		apiKey:  apiKey,
		baseURL: BASE_URL,
		client:  &http.Client{},
	}
}

func (adapter *PaymentRequestAdapter) CreatePaymentRequest(pr models.PaymentRequest) (models.CreatePaymentRequestResp, error) {
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

	resp, err := adapter.client.Do(req)
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

func (adapter *PaymentRequestAdapter) ListPaymentRequest(flag models.ListFlag) (models.ListPaymentRequestResp, error) {
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

	resp, err := adapter.client.Do(req)
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

func (adapter *PaymentRequestAdapter) FetchPaymentRequest(requestCode string) (models.FetchPaymentRequestResp, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", BASE_URL, PAYMENT_REQUEST_URL, requestCode), nil)

	if err != nil {
		return models.FetchPaymentRequestResp{}, err
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("PAYSTACK_SK"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := adapter.client.Do(req)
	if err != nil {
		return models.FetchPaymentRequestResp{}, err
	}
	defer resp.Body.Close()

	var body models.FetchPaymentRequestResp

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return models.FetchPaymentRequestResp{}, err
	}

	if !body.Status {
		return models.FetchPaymentRequestResp{}, errors.New(body.Message)
	}

	return body, nil
}
