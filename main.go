package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := os.Getenv("HTTP_HELLO_MSG")
	if message == "" {
		message = "Hello, World!"
	}

	fmt.Fprintf(w, message)
}

func main() {
	port := os.Getenv("HTTP_HELLO_PORT")
	if port == "" {
		port = "8080"
	}
	port = fmt.Sprintf(":%s", port)

	http.HandleFunc("/", helloHandler)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
