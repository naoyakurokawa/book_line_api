package handler

import (
	"context"

	"github.com/naoyakurokawa/book_line_api/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListBooksService RegisterUserService LoginService
type ListBooksService interface {
	ListBooks(ctx context.Context) (entity.Books, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
}

type LoginService interface {
	Login(ctx context.Context, name, pw string) (string, error)
}
