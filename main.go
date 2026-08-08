package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleHello)

	http.ListenAndServe(":8080", nil)
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
