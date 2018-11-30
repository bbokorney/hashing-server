package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// only allow POST requests
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("TODO implement handler")
}

func parsePortFromEnv() string {
	defaultPort := 8080
	envPort := os.Getenv("PORT")
	port := 0
	if envPort != "" {
		if i, err := strconv.Atoi(envPort); err == nil {
			port = i
		} else {
			log.Fatalf("Value for 'PORT' environment variable must be an int")
		}
	} else {
		log.Printf("Using default port %d", defaultPort)
		port = 8080
	}
	return fmt.Sprintf(":%d", port)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hash/", apiHandler{})
	addr := parsePortFromEnv()
	log.Printf("Starting server on %s", addr)
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Fatal(srv.ListenAndServe())
}
