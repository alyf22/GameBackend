package Handlers

import (
	"mygame/Services"
)

type SessionHandler struct {
	service        *Services.AuthService
	sessionService *Services.SessionService
}

func NewSessionHandler(authService *Services.AuthService, sessionService *Services.SessionService) *SessionHandler {
	return &SessionHandler{
		service:        authService,
		sessionService: sessionService,
	}
}

func (h *SessionHandler) CreateSession(userID string) (string, error) {
	token, err := h.sessionService.CreateToken(userID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (h *SessionHandler) ValidateSession(token string) (string, error) {
	userID, err := h.sessionService.ValidateToken(token)
	if err != nil {
		return "", err
	}
	return userID, nil
}
