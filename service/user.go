package service

import (
	"context"
	"fmt"

	"github.com/naoyakurokawa/book_line_api/entity"
	"github.com/naoyakurokawa/book_line_api/store"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserService struct {
	DB   store.Execer
	Repo UserCreator
}

func (cu *CreateUserService) CreateUser(ctx context.Context, name, password, role string) (*entity.User, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}
	u := &entity.User{
		Name:     name,
		Password: string(pw),
		Role:     role,
	}

	if err := cu.Repo.CreateUser(ctx, cu.DB, u); err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	return u, nil
}
