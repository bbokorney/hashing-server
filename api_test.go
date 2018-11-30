package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiGet(t *testing.T) {
	ts := httptest.NewServer(apiHandler{})
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("Expected status MethodNotAllowed for GET but received %s", res.Status)
	}
}
