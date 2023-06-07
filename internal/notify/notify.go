package notify

import (
	"log"
	"net/http"
)

func SendFailure(target string) (*http.Response, error) {
	// TODO refactor to change http methods, add auth, add body/payload
	res, err := http.Get(target)
	if err != nil {
		log.Fatalf("Http GET failed")
	}
	defer res.Body.Close()
	// TODO create error path for retry
	return res, nil
}

func SendSuccss(target string) error {
	log.Panicf("SendSuccss not yet implemented")
	return nil
}
