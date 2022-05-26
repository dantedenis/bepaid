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
		auth:    base64.StdEncoding.EncodeToString([]byte(username + ":" + password))}
}

func (a *Api) Authorizations(ctx context.Context, request *vo.AuthorizationRequest) (*http.Response, error) {
	return a.sendRequest(ctx, http.MethodPost, authorizations, &request)
}

func (a *Api) Captures(ctx context.Context, request *vo.CaptureRequest) (*http.Response, error) {
	return a.sendRequest(ctx, http.MethodPost, captures, &request)
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
	r.Header.Set("Authorization", "Basic "+a.auth)

	if method == http.MethodPost {
		//r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		r.Header.Set("Content-Type", "application/json")
	}

	r.Header.Set("Accept", "application/json")

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
