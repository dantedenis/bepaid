package service

import (
	"bepaid-sdk/api/contracts"
	"bepaid-sdk/service/vo"
	"context"
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
		 "request": map[string]interface{} {
		    "parent_uid": captureRequest.Parent_uid,
		    "amount": captureRequest.Amount,
		  },
	}
	bytesRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := NewRequestWithContext(ctx, "POST", vo.URLRequest, bytes.NewBuffer(bytesRequest))
	if err != nil {
		return nil, err
	}
	defer resb.Body.close()
	
}
