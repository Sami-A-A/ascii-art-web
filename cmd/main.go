package main

import (
	"asciiartweb/routes"
	"net/http"
	"log"
	"os"
)

func main() {

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Creates configurable Serve Mux
	serveMux := http.NewServeMux()

	// Initialize Routes
	routes.UseRoutes(serveMux)

	// File Servers
	static := http.FileServer(http.Dir("../static"))
	styles := http.FileServer(http.Dir("../styles"))

	serveMux.Handle("/static/assets/", http.StripPrefix("/static/assets/", static))
	serveMux.Handle("/static/asciifonts/", http.StripPrefix("/static/asciifonts/", static))
	serveMux.Handle("/styles/", http.StripPrefix("/styles", styles))

	// Port and Serve Mux with error handling for running server
	err := http.ListenAndServe(":" + port, serveMux)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	} else {
		log.Printf("Server is listening on port %s\n", port)
	}

}
