package user

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct{ DB *sqlx.DB }

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{DB: db}
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.DB.Get(&user, "SELECT id, email, password FROM users WHERE email = ? LIMIT 1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(u *User) error {
	query := "INSERT INTO users (email, password) VALUES (:email, :password)"
	_, err := r.DB.NamedExec(query, u)
	return err
}
