package handler

import (
	"context"

	"github.com/naoyakurokawa/book_line_api/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListBooksService
type ListBooksService interface {
	ListBooks(ctx context.Context) (entity.Books, error)
}
