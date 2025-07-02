package user

import (
	"Test_Go/model"
	"Test_Go/repository"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{ repo repository.UserRepository }

func NewUserService(r repository.UserRepository) UserService { return UserService{repo: r} }

func (s *UserService) Register(email, password string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := model.User{
		Email:    sql.NullString{String: email, Valid: true},
		Password: sql.NullString{String: string(hash), Valid: true},
	}
	return s.repo.Create(user)
}

func (s *UserService) Login(email, password string) (*model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user.ID.Int64 == 0 {
		return nil, errors.New("user not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
