package Services

import (
	"errors"
	"mygame/Repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AuthRepository *Repositories.AuthRepository
}

func NewAuthService(authRepo *Repositories.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository: authRepo,
	}
}

func (s *AuthService) Authenticate(email, password string) (string, error) {
	user, err := s.AuthRepository.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return "dummy_token", nil
}
