// package main

// import (
// 	"log"
// )

// func main() {

// 	store, err := NewPostgresStore()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := store.Init(); err != nil {
// 		log.Fatal(err)
// 	}
// 	server := NewAPIServer(":3000", store)
// 	server.Run()

// }

// main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	hello()
	// Define a handler function for handling incoming requests
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}

	// Register the handler function for the root ("/") path
	http.HandleFunc("/", handler)

	// Start the server and listen on port 8080
	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
