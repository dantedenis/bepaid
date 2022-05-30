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
	defer ctrl.Finish()

	capture := testdata.NewMockApi(ctrl)
	capture.EXPECT().Capture(context.Background(), vo.NewCaptureRequest(50, "1-310b0da80b")).Return(http.Response{StatusCode: 200, Body: ""})

	captureTest := NewApiService(capture)
	//todo проверка результатов
	response, err := captureTest.Capture(context.Background(), *vo.NewCaptureRequest(50, "1-310b0da80b"))

	if err != nil {
		t.Error(err)
	}

}
