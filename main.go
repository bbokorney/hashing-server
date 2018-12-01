package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
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
	// handle the SIGINT OS signal and initiate
	// a graceful shutdown
	var wg sync.WaitGroup
	shutdownChan := make(chan bool, 1)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		log.Println("Starting shutdown")
		// we've received SIGINT, close
		// the shutdownChan to indicate the
		// server is shutting down soon
		close(shutdownChan)
		// wait for any existing requests to complete
		wg.Wait()
		// exit the process
		os.Exit(0)
	}()

	mux := http.NewServeMux()
	mux.Handle("/hash/", apiHandler{
		shutdownChan: shutdownChan,
		wg:           &wg,
	})

	addr := parsePortFromEnv()
	log.Printf("Starting server on %s", addr)
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Fatal(srv.ListenAndServe())
}
