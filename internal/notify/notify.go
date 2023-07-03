package notify

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

func HttpReq(method string, target string, json_body []byte) (*http.Response, error) {
	// TODO add auth headers
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
	req.Header = http.Header{
		"Accept":       {"application/json"},
		"Content-Type": {"application/json"},
	}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("HTTP request failed, retrying")
		try := 3
		count := 0
		for {
			count += 1
			log.Printf("Retry %v/%v", count, try)
			res, err = client.Do(req)
			if err == nil {
				break
			}
			if count >= try {
				return nil, fmt.Errorf("HTTP request failed with error: %w", err)
			}
			time.Sleep(3 * time.Second)
		}

	}
	defer res.Body.Close()
	return res, nil
}
