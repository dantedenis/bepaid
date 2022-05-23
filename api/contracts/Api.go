package contracts

import (
	"bepaid-sdk/service/vo"
	"context"
	"net/http"
)

type Api interface {
	Authorization(ctx context.Context, request vo.AuthorizationRequest) (http.Response, error)
	Capture(ctx context.Context, capture vo.CaptureRequest) (http.Response, error)
}
