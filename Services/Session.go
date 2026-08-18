package Services

import (
	"mygame/Repositories"
)

type SessionService struct {
	SessionRepository *Repositories.SessionRepository
	TokenService      *TokenService
}

func NewSessionService(sessionRepo *Repositories.SessionRepository, tokenService *TokenService) *SessionService {
	return &SessionService{
		SessionRepository: sessionRepo,
		TokenService:      tokenService,
	}
}

func (s *SessionService) CreateToken(userID string) (string, error) {
	token, err := s.TokenService.GenerateToken(userID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *SessionService) ValidateToken(token string) (string, error) {
	userID, err := s.TokenService.ValidateToken(token)
	if err != nil {
		return "", err
	}
	return userID, nil
}
