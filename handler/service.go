package handler

import (
	"context"

	"github.com/naoyakurokawa/book_line_api/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . FetchBooksService FetchBookMemosService CreateUserService LoginService

// type BookServicer interface {
// 	CreateBook(ctx context.Context, isbn int64) error
// 	FetchBooks(ctx context.Context) (entity.Books, error)
// }

type CreateBookServicer interface {
	CreateBook(ctx context.Context, isbn int64) error
}

type FetchBooksServicer interface {
	FetchBooks(ctx context.Context) (entity.Books, error)
}

type CreateBookMemoServicer interface {
	CreateBookMemo(ctx context.Context, book_ID entity.BookID, page int64, detail string) error
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
