package hackernews

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

// This file contains helpers for the tests in this package.

// Load test data
func helperLoadBytes(t *testing.T, name string) []byte {
	path := filepath.Join("./testdata", name) // relative path
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}

func mockServerHelper(mockResponse []byte) (*Client, *httptest.Server) {
	// setup mock client
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(mockResponse)
	}))

	client := &Client{
		http:   server.Client(),
		apiURL: server.URL,
	}

	return client, server
}
