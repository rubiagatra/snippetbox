package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippet Box"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fmt.Println("Hello Snippet Box")

	log.Println("Starting server on:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
