package api

import (
	"bepaid-sdk/service/vo"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

const (
	payments       = "payments"
	authorizations = "authorizations"
	captures       = "captures"
	voids          = "voids"
	refunds        = "refunds"
)

type Api struct {
	client  *http.Client
	baseUrl string
	auth    string
}

func NewApi(client *http.Client, baseUrl, username, password string) *Api {
	return &Api{
		client:  client,
		baseUrl: baseUrl,
		auth:    "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))}
}

func (a *Api) Payments(ctx context.Context, payment vo.PaymentRequest) (*http.Response, error) {
	return a.sendRequest(ctx, http.MethodPost, payments, &payment)
}

func (a *Api) Authorizations(ctx context.Context, authorization vo.AuthorizationRequest) (*http.Response, error) {
	return a.sendRequest(ctx, http.MethodPost, authorizations, &authorization)
}

func (a *Api) Captures(ctx context.Context, capture vo.CaptureRequest) (*http.Response, error) {
	return a.sendRequest(ctx, http.MethodPost, captures, &capture)
}

func (a *Api) Voids(ctx context.Context, void vo.VoidRequest) (*http.Response, error) {
	return a.sendRequest(ctx, http.MethodPost, voids, &void)
}

func (a *Api) Refunds(ctx context.Context, refund vo.RefundRequest) (*http.Response, error) {
	return a.sendRequest(ctx, http.MethodPost, refunds, &refund)
}

func (a *Api) sendRequest(ctx context.Context, method, path string, request interface{}) (*http.Response, error) {

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequestWithContext(ctx, method, a.baseUrl+path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	r.Header.Set("Authorization", a.auth)
	r.Header.Set("Accept", "application/json")

	if method == http.MethodPost {
		//r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r.Header.Set("Content-Type", "application/json")
	}

	return a.client.Do(r)
}

// if request doesnt have "request" field
func marshalRequest(request interface{}) (io.Reader, error) {
	b, err := json.Marshal(struct {
		Request interface{} `json:"request"`
	}{request})
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
