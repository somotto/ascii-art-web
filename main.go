package main

import (
	"ascii/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
    // Register the handler for the home page
    http.HandleFunc("/", handlers.HomeHandler)
    
    // Register the handler for the ASCII art generation
    http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

    fmt.Println("Server is running on http://localhost:8080")
    // Start the HTTP server and log any errors
    log.Fatal(http.ListenAndServe(":8080", nil))
}