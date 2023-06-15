package notify

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var httpreqvalues = []struct {
	method string
	body   []byte
}{
	{"GET", []byte(`{"test": "value"}`)},
	{"GET", nil},
	{"PUT", []byte(`{"test": "value"}`)},
	{"PUT", nil},
	{"POST", []byte(`{"test": "value"}`)},
	{"POST", nil},
}

func TestHttpReq(t *testing.T) {
	for _, tt := range httpreqvalues {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			want := tt.method
			if r.Method != want {
				t.Errorf("Wrong method, is %v instead of %v", r.Method, want)
			}
		}))

		res, err := HttpReq(tt.method, ts.URL, tt.body)
		// TODO read about json in golang: https://go.dev/blog/json
		// Find best way to pass body to functions and support strings/json/etc
		if err != nil {
			t.Errorf("Error in SendFailure request: %v", err)
		}
		if res.StatusCode != 200 {
			t.Errorf("Return status was not 200. Status code: %v", res.StatusCode)
		}
		ts.Close()
	}
}

var httpreqfailvalues = []struct {
	method string
	body   []byte
}{
	{"LOCK", nil},
}

func TestHttpReqFail(t *testing.T) {
	for _, tt := range httpreqfailvalues {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		}))

		_, err := HttpReq(tt.method, ts.URL, tt.body)

		if err == nil {
			t.Log(err)
			t.Errorf("HTTP request should have failed but did not. Test values were: %v, %v", tt.method, tt.body)
		}
		ts.Close()
	}
}
