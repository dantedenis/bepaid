package api

import (
	"bepaid-sdk/service/vo"
	"context"
	"net/http"
)

type Api struct {
	client  *http.Client
	baseUrl string
}

func NewApi(client *http.Client, baseUrl string) *Api {
	return &Api{client: client, baseUrl: baseUrl}
}

func SimpleCreate(url string) Api {
	return Api{&http.Client{}, url}
}

func (a Api) Authorization(ctx context.Context, request vo.AuthorizationRequest) (http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (a Api) Capture(ctx context.Context, capture vo.CaptureRequest) (http.Response, error) {
	//TODO implement me
	panic("implement me")
}
