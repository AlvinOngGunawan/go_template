package repository

import (
	"Test_Go/model"
	"database/sql"
	"errors"
)

type UserRepository struct{ DB *sql.DB }

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{DB: db}
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	query := "SELECT id, email, password FROM users WHERE email = ? LIMIT 1"

	var user model.User

	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // no user found
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(u model.User) error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`

	_, err := r.DB.Exec(query, u.Email.String, u.Password.String)
	return err
}
