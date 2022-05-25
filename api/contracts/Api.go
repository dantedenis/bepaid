package contracts

import (
	"bepaid-sdk/service/vo"
	"context"
	"net/http"
)

//go:generate mockgen -source=Api.go -destination=../../testdata/ApiMock.go -package=testdata
type Api interface {
	Authorizations(ctx context.Context, request *vo.AuthorizationRequest) (*http.Response, error)
	Captures(ctx context.Context, capture *vo.CaptureRequest) (*http.Response, error)
}
