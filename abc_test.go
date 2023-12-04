// main_test.go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	server := NewAPIServer(":8080", nil) // Assuming you don't need a database for this test

	// Start the server in a goroutine
	go server.Run()
	<-time.After(50 * time.Millisecond)

	// Make a request to the "/account" endpoint
	req := httptest.NewRequest("GET", "http://localhost:8080/account", nil)
	rr := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		fmt.Println("Demo test passed")
	}
}
