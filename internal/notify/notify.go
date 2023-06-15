package notify

import (
	"bytes"
	"fmt"
	"net/http"
)

func HttpReq(method string, target string, json_body []byte) (*http.Response, error) {
	// TODO refactor to change http methods, add auth, add body/payload
	validmethods := map[string]bool{
		"GET":  true,
		"PUT":  true,
		"POST": true,
	}
	if _, valid := validmethods[method]; !valid {
		return nil, fmt.Errorf("HTTP request method invalid, tried: %v", method)
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, target, bytes.NewBuffer(json_body))
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed with error: %w", err)
	}
	defer res.Body.Close()
	// TODO create error path for retry
	return res, nil
}
