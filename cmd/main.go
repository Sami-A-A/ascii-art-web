package main

import (
	"asciiartweb/routes"
	"net/http"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Creates configurable Serve Mux
	serveMux := http.NewServeMux()

	// Initialize Routes
	routes.UseRoutes()

	// Port and Serve Mux with error handling for server
	err := http.ListenAndServe(":" + port, serveMux)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	} else {
		log.Printf("Server is listening on port %s\n", port)
	}

}

// func main() {
// 	// Define the handler function for the "/" route
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "Hello, World!")
// 	})

// 	// Start the server on port 8080
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		log.Fatalf("Error starting server: %s\n", err)
// 	}
// }