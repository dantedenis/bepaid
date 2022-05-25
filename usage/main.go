package main

import (
	"bepaid-sdk/api"
	"bepaid-sdk/service"
	"bepaid-sdk/service/vo"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {

	baseURL := "https://gateway.bepaid.by/transactions/"
	shopId := ""
	secret := ""

	api1 := api.NewApi(http.DefaultClient, api.DefaultEndpoints, baseURL, shopId, secret)

	service1 := service.NewApiService(api1)

	tr, err := service1.Authorizations(context.Background(), vo.NewAuthorizationRequest())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tr)
}
