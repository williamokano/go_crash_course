package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":3000", nil)
}
