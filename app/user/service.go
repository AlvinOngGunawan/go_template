package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{ repo UserRepository }

func NewUserService(r UserRepository) UserService { return UserService{repo: r} }

func (s *UserService) Register(email, password string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &User{
		Email:    email,
		Password: string(hash),
	}
	return s.repo.Create(user)
}

func (s *UserService) Login(email, password string) (*User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
