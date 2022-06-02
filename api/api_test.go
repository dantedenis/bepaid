package api

import (
	"bepaid-sdk/service/vo"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"testing"
	"time"
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

	// TODO remove secrets
	api2 = NewApi(http.DefaultClient, "https://gateway.bepaid.by", "", "")
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
					return api.Payment(context.TODO(), tc.req)
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
					return api.Authorization(context.TODO(), tc.req)
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
					return api.Capture(context.TODO(), tc.req)
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
					return api.Void(context.TODO(), tc.req)
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
					return api.Refund(context.TODO(), tc.req)
				})
		})
	}
}

func TestApi_Payment(t *testing.T) {
	cc := vo.NewCreditCard("4200000000000000", "123", "tim", "01", "2024")
	r := vo.NewPaymentRequest(int64(100), "RUB", "it's description", "mytrackingid", true, *cc).WithDuplicateCheck(false)

	resp, err := api2.Payment(context.Background(), *r)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Logf("resp.Body: %s", string(b))
		t.Fatalf("unexpected status code: %v", resp.StatusCode)
	}
	if len(b) == 0 {
		t.Fatal("Response body length == 0")
	}

	t.Logf("resp.Body: %s", string(b))
}

func TestApi_Authorization(t *testing.T) {
	cc := vo.NewCreditCard("4200000000000000", "123", "tim", "01", "2024")
	r := vo.NewAuthorizationRequest(int64(100), "RUB", "it's description", "mytrackingid", true, *cc).WithDuplicateCheck(false)

	resp, err := api2.Authorization(context.Background(), *r)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Logf("resp.Body:\n%s", string(b))
		t.Fatalf("unexpected status code: %v", resp.StatusCode)
	}
	if len(b) == 0 {
		t.Fatal("Response body length == 0")
	}

	t.Logf("resp.Body:\n%s", string(b))

}

func TestApi_AuthorizationCapture(t *testing.T) {
	amount := rand.New(rand.NewSource(time.Now().Unix())).Int63() % 100

	cc := vo.NewCreditCard("4200000000000000", "123", "tim", "01", "2024")
	r := vo.NewAuthorizationRequest(amount, "RUB", "it's description", "mytrackingid", true, *cc)

	resp, err := api2.Authorization(context.Background(), *r)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	buf := bytes.Buffer{}
	teeReader := io.TeeReader(resp.Body, &buf)
	uid := getUid(t, teeReader)

	cr := vo.NewCaptureRequest(uid, amount)

	resp, err = api2.Capture(context.Background(), *cr)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	buf = bytes.Buffer{}
	teeReader = io.TeeReader(resp.Body, &buf)
	uid = getUid(t, teeReader)
}

func TestApi_AuthorizationVoid(t *testing.T) {
	amount := rand.New(rand.NewSource(time.Now().Unix())).Int63() % 100

	cc := vo.NewCreditCard("4200000000000000", "123", "tim", "01", "2024")
	r := vo.NewAuthorizationRequest(amount, "RUB", "it's description", "mytrackingid", true, *cc)

	resp, err := api2.Authorization(context.Background(), *r)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	buf := bytes.Buffer{}
	teeReader := io.TeeReader(resp.Body, &buf)
	uid := getUid(t, teeReader)

	vr := vo.NewVoidRequest(uid, amount)

	resp, err = api2.Void(context.Background(), *vr)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	buf = bytes.Buffer{}
	teeReader = io.TeeReader(resp.Body, &buf)
	uid = getUid(t, teeReader)
}

func TestApi_AuthorizationCaptureRefund(t *testing.T) {
	amount := rand.New(rand.NewSource(time.Now().Unix())).Int63() % 100

	cc := vo.NewCreditCard("4200000000000000", "123", "tim", "01", "2024")
	r := vo.NewAuthorizationRequest(amount, "RUB", "it's description", "mytrackingid", true, *cc)

	resp, err := api2.Authorization(context.Background(), *r)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	buf := bytes.Buffer{}
	teeReader := io.TeeReader(resp.Body, &buf)
	uid := getUid(t, teeReader)

	cr := vo.NewCaptureRequest(uid, amount)

	resp, err = api2.Capture(context.Background(), *cr)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	buf = bytes.Buffer{}
	teeReader = io.TeeReader(resp.Body, &buf)
	uid = getUid(t, teeReader)

	rr := vo.NewRefundRequest(uid, amount, "need my money back")

	resp, err = api2.Refund(context.Background(), *rr)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	buf = bytes.Buffer{}
	teeReader = io.TeeReader(resp.Body, &buf)
	uid = getUid(t, teeReader)

}

func TestApi_PaymentRefund(t *testing.T) {
	r := vo.NewCaptureRequest("151281134-8d2c74c539", int64(100)).WithDuplicateCheck(false)

	resp, err := api2.Capture(context.Background(), *r)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Logf("resp.Body:\n%s", string(b))
		t.Fatalf("unexpected status code: %v", resp.StatusCode)
	}
	if len(b) == 0 {
		t.Fatal("Response body length == 0")
	}

	t.Logf("resp.Body:\n%s", string(b))

}

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

func getUid(t *testing.T, body io.Reader) string {
	m := map[string]interface{}{}

	err := json.NewDecoder(body).Decode(&m)
	if err != nil {
		t.Fatalf("Decoder.Decode: err is not nil: %v", err)
	}

	if len(m) == 0 {
		t.Fatal("Response body length == 0")
	}

	transactionMap, ok := m["transaction"]
	if !ok {
		t.Fatal("No 'transaction' key in map")
	}

	uid, ok := transactionMap.(map[string]interface{})["uid"]
	if !ok {
		t.Fatal("No 'uid' key in transactionMap")
	}

	uidS, ok := uid.(string)
	if !ok {
		t.Fatal("Value in 'uid' key is not a string")
	}

	return uidS
}
