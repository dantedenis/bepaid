package contracts

import (
	"bepaid-sdk/service/vo"
	"context"
)

//go:generate mockgen -source=service.go -destination=../../testdata/ServiceMock.go -package=testdata
type ApiService interface {
	Authorizations(ctx context.Context, authorizationRequest vo.AuthorizationRequest) (vo.TransactionResponse, error)
	Capture(ctx context.Context, captureRequest vo.CaptureRequest) (vo.TransactionResponse, error)
}
