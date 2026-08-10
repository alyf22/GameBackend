package Repositories

import (
	"context"

	"mygame/Models"

	"github.com/jackc/pgx/v5"
)

type AuthRepository struct {
	db *pgx.Conn
}

func NewAuthRepository(db *pgx.Conn) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) GetUserByEmail(email string) (*Models.User, error) {
	var user Models.User
	query := "SELECT id, name, email, password_hash FROM users WHERE email=$1"
	err := r.db.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
