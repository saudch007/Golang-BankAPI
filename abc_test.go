// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	store, err := NewPostgresStore()
	if err != nil {
		t.Fatalf("Error creating Postgres store: %v", err)
	}
	defer store.db.Close()

	if err := store.Init(); err != nil {
		t.Fatalf("Error initializing Postgres store: %v", err)
	}

	server := NewAPIServer(":8081", store)

	// Start the server in a goroutine
	go server.Run()
	<-time.After(50 * time.Millisecond)

	// Make a request to the "/account" endpoint
	req := httptest.NewRequest("GET", "http://localhost:8081/account", nil)
	rr := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", status, http.StatusOK)
	}
}
