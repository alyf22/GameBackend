package handlers

import (
	"encoding/json"
	"mygame/Models"
	"mygame/Services"
	"net/http"
)

type AuthHandler struct {
	service *Services.AuthService
}

func NewAuthHandler(service *Services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req Models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.service.Authenticate(
		req.Email,
		req.Password,
	)

	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	response := Models.LoginResponse{
		Token: token,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err = h.service.Register(
		req.Name,
		req.Email,
		req.Password,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
