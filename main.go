package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloServer)
	log.Println("Server started on 8080..")
	http.ListenAndServe(":8080", nil)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, fmt.Sprintf("Hello, %s! Version 2", name))
	log.Println(fmt.Sprintf("Hello, %s! Version 2", name))
}
