package contracts

import (
	"bepaid-sdk/service/vo"
	"context"
	"net/http"
)

//go:generate mockgen -source=Api.go -destination=../../testdata/ApiMock.go -package=testdata
type Api interface {
	Payments(ctx context.Context, payment vo.PaymentRequest) (*http.Response, error)
	Authorizations(ctx context.Context, authorization vo.AuthorizationRequest) (*http.Response, error)
	Captures(ctx context.Context, capture vo.CaptureRequest) (*http.Response, error)
	Voids(ctx context.Context, void vo.VoidRequest) (*http.Response, error)
	Refunds(ctx context.Context, refund vo.RefundRequest) (*http.Response, error)
}
