package service

import (
	"bepaid-sdk/service/vo"
	"bepaid-sdk/testdata"
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestApiService_Authorizations(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	api := testdata.NewMockApi(ctrl)
	api.EXPECT().Authorization(context.Background(), vo.NewAuthorizationRequest()).Return(http.Response{StatusCode: 200})

	//service := NewApiService(api)
	//response, error := service.Authorizations(context.Background(), vo.AuthorizationRequest{})
}

func TestApiService_Capture(t *testing.T) {
	ctrl := gomock.NewController(t)

}
