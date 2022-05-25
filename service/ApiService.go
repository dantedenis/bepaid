package service

import (
	"bepaid-sdk/api/contracts"
	"bepaid-sdk/service/vo"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ApiService struct {
	api contracts.Api
}

func NewApiService(api contracts.Api) *ApiService {
	return &ApiService{api: api}
}

func (a ApiService) Authorizations(ctx context.Context, authorizationRequest vo.AuthorizationRequest) (vo.TransactionResponse, error) {
	//TODO implement me
	panic("implement me")

}

func (a ApiService) Capture(ctx context.Context, captureRequest vo.CaptureRequest) (vo.TransactionResponse, error) {
	request := map[string]interface{}{
		"request": map[string]interface{}{
			"parent_uid": captureRequest.ParentUid,
			"amount":     captureRequest.Amount,
		},
	}
	bytesRequest, err := json.Marshal(request)
	if err != nil {
		return vo.TransactionResponse{}, err
	}
	resp, err := http.NewRequestWithContext(ctx, "POST", vo.URLRequest, bytes.NewBuffer(bytesRequest))
	if err != nil {
		return vo.TransactionResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)
	var result vo.TransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return vo.TransactionResponse{}, err
	}
	return result, nil
}
