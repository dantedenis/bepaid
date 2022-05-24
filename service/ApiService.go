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
	ctx, cancelCtx = context.WithTimeout(ctx, 3 * time.Second)
	respChan, errChan := make(chan *Response), make(chan error)
	request := map[string]interface{} /// -> TODO: roadmap
	
	go func(){
		bytes, err := json.Marshal(request)
		if err != nil {
			errChan <- err
		}
		resp, err := http.POST('https://gateway.bepaid.by/transactions/captures','application/json', bytes.NewBuffer(bytes))
		if err != nil {
			errChan <- err
		}
		respChan <- resp
	}()
	select {
	case <- cts.Done():
		cancelCtx()
	case <- respChan:
		
	case <- errChan:
		
	}
}
