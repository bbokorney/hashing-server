package main

import (
	"log"
	"net/http"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

	w.Write([]byte(Hash(password)))
}
