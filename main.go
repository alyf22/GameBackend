package main

import (
	"mygame/Database"
	Handlers "mygame/Handlers"
	"mygame/Repositories"
	"mygame/Services"
	"net/http"
	//"encoding/json"
)

func main() {
	db, err := Database.Connect()
	if err != nil {
		panic(err)
	}
	authRepo := Repositories.NewAuthRepository(db)
	authService := Services.NewAuthService(authRepo)
	authHandler := Handlers.NewAuthHandler(authService)
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleHello)
	mux.HandleFunc("/login", authHandler.Login)

	http.ListenAndServe(":8080", nil)
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
