package Services

import (
	"crypto/rand"
	"encoding/base64"
)

type TokenService struct{}

func (ts *TokenService) GenerateToken(userID string) (string, error) {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

func (ts *TokenService) ValidateToken(token string) (string, error) {
	// Implement token validation logic here
	// For example, you can use JWT to validate the token and extract the user ID
	return "", nil
}
