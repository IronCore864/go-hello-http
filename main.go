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
		name = "v0.0.16"
	}
	log.Println(fmt.Sprintf("Hello, %s!", name))
	
	w.Header().Set("Content-Type", "application/json")
	response := []byte(fmt.Sprintf("{\"test\": \"Hello, %s!\"}", name))
	_, err := w.Write(response)
	if err != nil {
        	// Handle the error gracefully
	}
}
