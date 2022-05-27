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

var ch = make(chan io.ReadCloser, 1)

type customRoundTripper struct{}

func (customRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	ch <- request.Body
	return nil, nil
}

//---------MarshalRequest---------
func TestApi_PaymentsMarshalRequest(t *testing.T) {

	api := NewApi(
		&http.Client{Transport: customRoundTripper{}},
		"",
		"",
		"",
	)

	tests := []struct {
		name string
		data P
		er   string
	}{
		{"test1", P{}, `{"request":{"amount":0,"currency":"","description":"","tracking_id":"","test":false,"credit_card":{"number":"","verification_value":"","holder":"","exp_month":"","exp_year":"","skip_three_d_secure_verification":false}}}`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// ignore response and err
			go api.Payments(context.TODO(), tc.data)

			body := <-ch
			defer body.Close()

			b, err := io.ReadAll(body)

			if err != nil {
				fatalfExpectedActual(t, "ReadAll returned not nil value", nil, err)
			}

			if string(b) != tc.er {
				fatalfExpectedActual(t, "Strings aren't equal", tc.er, string(b))
			}
		})
	}
}

func TestApi_AuthorizationsMarshalRequest(t *testing.T) {

	api := NewApi(
		&http.Client{Transport: customRoundTripper{}},
		"",
		"",
		"",
	)

	tests := []struct {
		name string
		data A
		er   string
	}{
		{"test1", A{}, `{"request":{"amount":0,"currency":"","description":"","tracking_id":"","test":false,"credit_card":{"number":"","verification_value":"","holder":"","exp_month":"","exp_year":"","skip_three_d_secure_verification":false}}}`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// ignore response and err
			go api.Authorizations(context.TODO(), tc.data)

			body := <-ch
			defer body.Close()

			b, err := io.ReadAll(body)

			if err != nil {
				t.Fatalf("ReadAll returned not nil value: \nER: %v,\nAR: %v", nil, err)
			}

			if string(b) != tc.er {
				t.Fatalf("wrong value: \nER: %v,\nAR: %v", tc.er, string(b))
			}
		})
	}
}

func TestApi_CapturesMarshalRequest(t *testing.T) {

	api := NewApi(
		&http.Client{Transport: customRoundTripper{}},
		"",
		"",
		"",
	)

	tests := []struct {
		name string
		data C
		er   string
	}{
		{"test1", C{}, `{"request":{"parent_uid":"","amount":0}}`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// ignore response and err
			go api.Captures(context.TODO(), tc.data)

			body := <-ch
			defer body.Close()

			b, err := io.ReadAll(body)

			if err != nil {
				t.Fatalf("ReadAll returned not nil value: \nER: %v,\n AR: %v", nil, err)
			}

			if string(b) != tc.er {
				t.Fatalf("wrong value: \nER: %v,\n AR: %v", tc.er, string(b))
			}
		})
	}
}

func TestApi_VoidsMarshalRequest(t *testing.T) {

	api := NewApi(
		&http.Client{Transport: customRoundTripper{}},
		"",
		"",
		"",
	)

	tests := []struct {
		name string
		data V
		er   string
	}{
		{"test1", V{}, `{"request":{"parent_uid":"","amount":0}}`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// ignore response and err
			go api.Voids(context.TODO(), tc.data)

			body := <-ch
			defer body.Close()

			b, err := io.ReadAll(body)

			if err != nil {
				t.Fatalf("ReadAll returned not nil value: \nER: %v,\n AR: %v", nil, err)
			}

			if string(b) != tc.er {
				t.Fatalf("wrong value: \nER: %v,\n AR: %v", tc.er, string(b))
			}
		})
	}
}

func TestApi_RefundsMarshalRequest(t *testing.T) {

	api := NewApi(
		&http.Client{Transport: customRoundTripper{}},
		"",
		"",
		"",
	)

	tests := []struct {
		name string
		data R
		er   string
	}{
		{"test1", R{}, `{"request":{"parent_uid":"","amount":0,"reason":""}}`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// ignore response and err
			go api.Refunds(context.TODO(), tc.data)

			body := <-ch
			defer body.Close()

			b, err := io.ReadAll(body)

			if err != nil {
				fatalfExpectedActual(t, "ReadAll returned not nil value", nil, err)
			}

			if string(b) != tc.er {
				fatalfExpectedActual(t, "String are not equal", tc.er, string(b))
			}
		})
	}
}

func fatalfExpectedActual(t *testing.T, msg string, er, ar interface{}) {
	t.Fatalf("%s:\nER: %v,\nAR: %v", msg, er, ar)
}
