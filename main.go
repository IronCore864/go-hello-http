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
		name = "v0.13.0"
	}
	log.Println(fmt.Sprintf("Hello, %s!", name))
	
	w.Header().Set("Content-Type", "application/json")
	response := []byte("{\"test\": \"Hello, world!\"}")
	_, err := w.Write(response)
	if err != nil {
        	// Handle the error gracefully
	}
}
