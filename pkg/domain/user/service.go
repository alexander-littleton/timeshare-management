package user

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user User) error
}

type Service struct {
	repo UserRepo
}

func NewService(repo UserRepo) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateUser(ctx context.Context, user User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), 14)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %w", err)
	}

	user.PasswordHash = string(bytes)

	return s.repo.CreateUser(ctx, user)
}
