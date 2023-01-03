package service

import (
	"context"

	"github.com/naoyakurokawa/book_line_api/entity"
	"github.com/naoyakurokawa/book_line_api/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . BookCreator BookFetcher BookMemoFetcher UserCreator UserFetcher TokenGenerator

type BookCreator interface {
	CreateBook(ctx context.Context, db store.Execer, b *entity.Book) error
}

type BookFetcher interface {
	FetchBooks(ctx context.Context, db store.Queryer) (entity.Books, error)
}

type BookMemoCreator interface {
	CreateBookMemo(ctx context.Context, db store.Execer, bm *entity.BookMemo) error
}

type BookMemoFetcher interface {
	FetchBookMemos(ctx context.Context, db store.Queryer, bookID string) (entity.BookMemos, error)
}

type UserCreator interface {
	CreateUser(ctx context.Context, db store.Execer, u *entity.User) error
}

type UserFetcher interface {
	FetchUserByID(ctx context.Context, db store.Queryer, name string) (*entity.User, error)
}

type TokenGenerator interface {
	GenerateToken(ctx context.Context, u entity.User) ([]byte, error)
}
