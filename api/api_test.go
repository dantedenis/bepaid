package api

import (
	"bepaid-sdk/service/vo"
	"context"
	"io"
	"net/http"
	"testing"
)

type A = vo.AuthorizationRequest

var ch = make(chan io.ReadCloser, 1)

type customRoundTripper struct{}

func (c customRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	ch <- request.Body
	return nil, nil
}

// TestApi_Authorizations asserts that vo.AuthorizationRequest marshaled correctly
func TestApi_AuthorizationsMarshalRequest(t *testing.T) {

	api := NewApi(
		&http.Client{Transport: customRoundTripper{}},
		",",
		"",
		"",
	)

	tests := []struct {
		name string
		data A
		er   string
	}{
		{"test1", A{}, `{"request":{"amount":0,"currency":"","tracking_id":"","test":false,"credit_card":{"number":"","verification_value":"","holder":"","exp_month":"","exp_year":"","skip_three_d_secure_verification":false}}}`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// ignore response nad err
			go api.Authorizations(context.TODO(), &tc.data)

			body := <-ch
			defer body.Close()

			b, err := io.ReadAll(body)

			if err != nil {
				t.Fatalf("ReadAll retrun not nil value: \nER: %v,\n AR: %v", nil, err)
			}

			if string(b) != tc.er {
				t.Fatalf("wrong value: \nER: %v,\n AR: %v", tc.er, string(b))
			}
		})
	}

}
