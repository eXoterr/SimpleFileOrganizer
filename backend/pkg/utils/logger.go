package utils

import (
	"log"
	"net/http"
)

func MustCheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func LogCheckError(err error) {
	if err != nil {
		log.Printf("error: %s", err)
	}

}

func HttpCheckError(err error, w http.ResponseWriter) {
	if err != nil {
		log.Printf("http error: %s", err)
	}

}

func LogRequest(r *http.Request) {
	log.Printf("request %s from %s", r.URL, r.RemoteAddr)
}
