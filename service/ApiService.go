package service

import (
	"bepaid-sdk/api/contracts"
	"bepaid-sdk/service/vo"
	"context"
	"encoding/json"
)

type ApiService struct {
	api contracts.Api
}

func NewApiService(api contracts.Api) *ApiService {
	return &ApiService{api: api}
}

func (c *ApiService) Authorizations(ctx context.Context, authorizationRequest vo.AuthorizationRequest) (*vo.TransactionResponse, error) {
	resp, err := c.api.Authorizations(ctx, authorizationRequest)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var tr vo.TransactionResponse
	if err = json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return nil, err
	}

	return &tr, err
}

func (c *ApiService) Capture(ctx context.Context, captureRequest vo.CaptureRequest) (*vo.TransactionResponse, error) {
	resp, err := c.api.Captures(ctx, captureRequest)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var tr vo.TransactionResponse
	if err = json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return nil, err
	}

	return &tr, err
}
