package handler

import (
	"context"

	"github.com/naoyakurokawa/book_line_api/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . FetchBooksService FetchBookMemosService CreateUserService LoginService
type FetchBooksService interface {
	FetchBooks(ctx context.Context) (entity.Books, error)
}

type FetchBookMemosService interface {
	FetchBookMemos(ctx context.Context, book_id string) (entity.BookMemos, error)
}

type CreateUserService interface {
	CreateUser(ctx context.Context, name, password, role string) (*entity.User, error)
}

type LoginService interface {
	Login(ctx context.Context, name, pw string) (string, error)
}
