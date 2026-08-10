package handlers

import (
	"encoding/json"
	"mygame/models"
	"mygame/services"
	"net/http"
)

type AuthHandler struct {
	service *services.AuthService
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest

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

	response := models.LoginResponse{
		Token: token,
	}

	json.NewEncoder(w).Encode(response)
}
