package main

import (
	"mygame/handlers"
	"net/http"
	//"encoding/json"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleHello)
	mux.HandleFunc("/login", (&handlers.AuthHandler{}).Login)

	http.ListenAndServe(":8080", nil)
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
