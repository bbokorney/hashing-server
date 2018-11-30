package main

import (
	"log"
	"net/http"
	"time"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// start a timer to ensure this entire request
	// takes at least 5 seconds
	timerChannel := time.After(5 * time.Second)

	// the timer isn't applied for errors

	// only allow POST requests to this route
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := req.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error parsing form body: %s", err)
		return
	}

	passwordArr := req.PostForm["password"]

	if passwordArr == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(passwordArr) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password := passwordArr[0]
	hashedPassword := Hash(password)

	// don't return until 5 seconds after receiving
	// this request has elasped
	<-timerChannel

	w.Write([]byte(hashedPassword))
}
