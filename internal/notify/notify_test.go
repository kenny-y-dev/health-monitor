package notify

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var failurevalues = []struct {
	method string
	body   []byte
}{
	{"GET", []byte(`{"test": "value"}`)},
	{"PUT", []byte(`{"test": "value"}`)},
	{"POST", []byte(`{"test": "value"}`)},
}

func TestFailure(t *testing.T) {
	//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		var want string
	//		if r.Method != want {
	//			t.Errorf("Wrong method, is %v instead of %v", r.Method, want)
	//		}
	//	}))
	//	defer ts.Close()

	for _, tt := range failurevalues {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			want := tt.method
			if r.Method != want {
				t.Errorf("Wrong method, is %v instead of %v", r.Method, want)
			}
		}))

		res, err := Failure(tt.method, ts.URL, tt.body)
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
