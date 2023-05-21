package notify

import (
	"log"
	"net/http"
)

func Send() {
	resp, err := http.Get("https://webhook.site/2e815d06-d941-4f1e-b74b-633c63509b4c")
	if err != nil {
		log.Fatalf("Http GET failed")
	}
	defer resp.Body.Close()
}
