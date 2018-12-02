package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func TestApiPost(t *testing.T) {
	ts := httptest.NewServer(apiHandler{})

	res, err := http.PostForm(ts.URL, url.Values{"password": {"angryMonkey"}})
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK but received %s", res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	bodyStr := string(body)

	base64Encoded := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
	if bodyStr != base64Encoded {
		t.Fatalf("Expected %s but got %s", base64Encoded, bodyStr)
	}
}

func TestApiBadRequest(t *testing.T) {
	ts := httptest.NewServer(apiHandler{})
	defer ts.Close()

	res, err := http.PostForm(ts.URL, url.Values{"foobar": {"angryMonkey"}})
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("Expected status BadRequest but received %s", res.Status)
	}
}
