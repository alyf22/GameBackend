package main

import (
	"encoding/json"
	"mygame/Database"
	"mygame/Handlers"
	"mygame/Repositories"
	"mygame/Services"
	"net/http"
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
	mux.HandleFunc("GET /", HandleHello)
	mux.HandleFunc("POST /login", authHandler.Login)
	mux.HandleFunc("POST /register", authHandler.Register)

	http.ListenAndServe(":8080", mux)
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
}
