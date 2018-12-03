package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func parsePortFromEnv() string {
	// this could be made simpler with a library like envconfig (https://github.com/kelseyhightower/envconfig)
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
	// Use shutdownInitiated to indicate when
	// to shut down the HTTP server
	shutdownInitiated := make(chan struct{})
	// shutdownCompleted will indicate when the server
	// shutdown has completed so we can exit the process
	shutdownCompleted := make(chan struct{})

	mux := http.NewServeMux()

	addr := parsePortFromEnv()
	log.Printf("Starting server on %s", addr)
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	var stats *statsHandler = &statsHandler{}

	mux.Handle("/hash/", statsWrapper(apiHandler{}, stats))
	mux.Handle("/stats/", stats)
	mux.HandleFunc("/shutdown/", func(w http.ResponseWriter, req *http.Request) {
		// we've received a request to shut down, close
		// the shutdownInitiated to indicate the
		// server should be shut down
		close(shutdownInitiated)
	})

	go func() {
		<-shutdownInitiated
		log.Println("Starting shutdown of server")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down server: %v", err)
		}
		log.Println("Server shut down")
		close(shutdownCompleted)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Error starting server: %v", err)
	}

	<-shutdownCompleted
}
