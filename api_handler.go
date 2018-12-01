package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type apiHandler struct {
	shutdownChan chan bool
	wg           *sync.WaitGroup
}

func shutdownInitiated(shutdownChan chan bool) bool {
	// No value will ever be sent on this channel.
	// If we can read from it, that's because it's been
	// closed. If we can't read from it, it hasn't been
	// closed, and the default case will be executed.
	select {
	case <-shutdownChan:
		return true
	default:
		return false
	}
}

func (a apiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// ensure a shutdown hasn't been initiated
	if shutdownInitiated(a.shutdownChan) {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	a.wg.Add(1)
	defer a.wg.Done()
	// start a timer to ensure this entire request
	// takes at least 5 seconds
	// the timer isn't applied for requests that result in errors
	timerChannel := time.After(5 * time.Second)

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
