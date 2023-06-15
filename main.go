package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloServer)
	log.Println("Server started on 8080.")
	http.ListenAndServe(":8080", nil)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "world"
	}
	log.Println(fmt.Sprintf("Hello, %s! 0.0.7", name))
	
	w.Header().Set("Content-Type", "application/json")
	response := []byte("Hello, world!")
	_, err := w.Write(response)
	if err != nil {
        	// Handle the error gracefully
	}
}
