package Handlers

import (
	"mygame/Services"
)

type SessionHandler struct {
	service *Services.AuthService
}

func NewSessionHandler(authService *Services.AuthService) *SessionHandler {
	return &SessionHandler{
		service: authService,
	}
}

func (h *SessionHandler) CreateSession(userID string) (string, error) {
	token, err := Services.GenerateToken(userID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (h *SessionHandler) ValidateSession(token string) (string, error) {
	userID, err := Services.ValidateToken(token)
	if err != nil {
		return "", err
	}
	return userID, nil
}
