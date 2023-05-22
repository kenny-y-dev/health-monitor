package notify

import (
	"log"
	"net/http"
)

func SendFailure(target string) error {
	// TODO refactor to change http methods, add auth, add body/payload
	resp, err := http.Get(target)
	if err != nil {
		log.Fatalf("Http GET failed")
	}
	defer resp.Body.Close()
	// TODO create error path for retry
	return nil
}

func SendSuccss(target string) error {
	log.Panicf("SendSuccss not yet implemented")
	return nil
}
