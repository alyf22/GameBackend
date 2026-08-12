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
	query := "SELECT id, name, email, password FROM users WHERE email=$1"
	err := r.db.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) CreateUser(user *Models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(context.Background(), query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}
