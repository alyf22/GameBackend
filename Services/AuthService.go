package services

type AuthService struct {
}

func (s *AuthService) Authenticate(email, password string) (string, error) {
	return "dummy_token", nil
}
