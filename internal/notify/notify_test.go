package notify

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendFailure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer ts.Close()

	res, err := SendFailure(ts.URL)
	if err != nil {
		t.Errorf("Error in SendFailure request: %v", err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Return status was not 200. Status code: %v", res.StatusCode)
	}
}
