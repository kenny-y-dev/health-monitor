package notify

import (
	"bytes"
	"log"
	"net/http"
)

func Failure(method string, target string, json_body []byte) (*http.Response, error) {
	// TODO refactor to change http methods, add auth, add body/payload
	client := &http.Client{}
	req, err := http.NewRequest(method, target, bytes.NewBuffer(json_body))
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Http request failed with error: %v", err)
	}
	defer res.Body.Close()
	// TODO create error path for retry
	return res, nil
}

func Succss(target string) error {
	log.Panicf("SendSuccss not yet implemented")
	return nil
}
