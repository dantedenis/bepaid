package api

import (
	"bepaid-sdk/service/vo"
	"context"
	"io"
	"net/http"
	"testing"
)

type P = vo.PaymentRequest
type A = vo.AuthorizationRequest
type C = vo.CaptureRequest
type V = vo.VoidRequest
type R = vo.RefundRequest

var (
	// All tests use this channel. So, no parallel testing.
	ch = make(chan io.ReadCloser, 1)

	// testingClient sends Request.Body to ch.
	testingClient = &http.Client{Transport: customRoundTripper{}}

	// Default Api used by all tests.
	api = Api{client: testingClient}
)

type customRoundTripper struct{}

func (customRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	ch <- request.Body
	return nil, nil
}

func TestApi_PaymentsMarshalRequest(t *testing.T) {

	tests := []struct {
		name string
		req  P
		er   string
	}{
		{"defaultValue", P{}, `{"request":{"amount":0,"currency":"","description":"","tracking_id":"","test":false,"credit_card":{"number":"","verification_value":"","holder":"","exp_month":"","exp_year":"","skip_three_d_secure_verification":false}}}`},
		{"requestConstructor", *vo.NewPaymentRequest(int64(1), "rub", "rub_1", "id1", true, *vo.NewCreditCard("5555", "123", "tim", "05", "2024")), `{"request":{"amount":1,"currency":"rub","description":"rub_1","tracking_id":"id1","test":true,"credit_card":{"number":"5555","verification_value":"123","holder":"tim","exp_month":"05","exp_year":"2024","skip_three_d_secure_verification":false}}}`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testMarshallRequest(
				t,
				tc.er,
				func() (*http.Response, error) {
					return api.Payments(context.TODO(), tc.req)
				})
		})
	}
}

func TestApi_AuthorizationsMarshalRequest(t *testing.T) {

	tests := []struct {
		name string
		req  A
		er   string
	}{
		{"defaultValue", A{}, `{"request":{"amount":0,"currency":"","description":"","tracking_id":"","test":false,"credit_card":{"number":"","verification_value":"","holder":"","exp_month":"","exp_year":"","skip_three_d_secure_verification":false}}}`},
		{"requestConstructor", *vo.NewAuthorizationRequest(int64(1), "rub", "rub_1", "id1", true, *vo.NewCreditCard("5555", "123", "tim", "05", "2024")), `{"request":{"amount":1,"currency":"rub","description":"rub_1","tracking_id":"id1","test":true,"credit_card":{"number":"5555","verification_value":"123","holder":"tim","exp_month":"05","exp_year":"2024","skip_three_d_secure_verification":false}}}`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testMarshallRequest(
				t,
				tc.er,
				func() (*http.Response, error) {
					return api.Authorizations(context.TODO(), tc.req)
				})
		})
	}
}

func TestApi_CapturesMarshalRequest(t *testing.T) {

	tests := []struct {
		name string
		req  vo.CaptureRequest
		er   string
	}{
		{"defaultValue", C{}, `{"request":{"parent_uid":"","amount":0}}`},
		{"requestConstructor", *vo.NewCaptureRequest("id123", int64(63)), `{"request":{"parent_uid":"id123","amount":63}}`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testMarshallRequest(
				t,
				tc.er,
				func() (*http.Response, error) {
					return api.Captures(context.TODO(), tc.req)
				})
		})
	}
}

func TestApi_VoidsMarshalRequest(t *testing.T) {

	tests := []struct {
		name string
		req  V
		er   string
	}{
		{"defaultValue", V{}, `{"request":{"parent_uid":"","amount":0}}`},
		{"requestConstructor", *vo.NewVoidRequest("id123", int64(63)), `{"request":{"parent_uid":"id123","amount":63}}`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testMarshallRequest(
				t,
				tc.er,
				func() (*http.Response, error) {
					return api.Voids(context.TODO(), tc.req)
				})
		})
	}
}

func TestApi_RefundsMarshalRequest(t *testing.T) {

	tests := []struct {
		name string
		req  R
		er   string
	}{
		{"defaultValue", R{}, `{"request":{"parent_uid":"","amount":0,"reason":""}}`},
		{"requestConstructor", *vo.NewRefundRequest("id123", int64(63), "reason"), `{"request":{"parent_uid":"id123","amount":63,"reason":"reason"}}`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testMarshallRequest(
				t,
				tc.er,
				func() (*http.Response, error) {
					return api.Refunds(context.TODO(), tc.req)
				})
		})
	}
}

//---------MarshalRequest---------

func testMarshallRequest(t *testing.T, er string, startRequest func() (*http.Response, error)) {
	// ignore response and error
	go startRequest()

	body := <-ch
	defer body.Close()

	b, err := io.ReadAll(body)

	if err != nil {
		fatalfWithExpectedActual(t, "ReadAll returned not nil value", nil, err)
	}

	if string(b) != er {
		fatalfWithExpectedActual(t, "Strings aren't equal", er, string(b))
	}
}

func fatalfWithExpectedActual(t *testing.T, msg string, er, ar interface{}) {
	t.Fatalf("%s:\nER: %v,\nAR: %v", msg, er, ar)
}
