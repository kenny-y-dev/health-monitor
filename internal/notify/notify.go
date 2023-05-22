package notify

import (
	"log"
	"net/http"
)

func Send(target string) {
	resp, err := http.Get(target)
	if err != nil {
		log.Fatalf("Http GET failed")
	}
	defer resp.Body.Close()
}
