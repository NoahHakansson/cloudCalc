package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// API routes
	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)
	port := ":3000"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}
