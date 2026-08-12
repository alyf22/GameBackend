package Services

import (
	"errors"
	"mygame/Models"
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return "dummy_token", nil
}

func (s *AuthService) Register(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &Models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.AuthRepository.CreateUser(user)
}
