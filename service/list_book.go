package service

import (
	"context"
	"fmt"

	"github.com/naoyakurokawa/book_line_api/entity"
	"github.com/naoyakurokawa/book_line_api/store"
)

type ListBook struct {
	DB   store.Queryer
	Repo BookLister
}

func (l *ListBook) ListBooks(ctx context.Context) (entity.Books, error) {
	bs, err := l.Repo.ListBooks(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return bs, nil
}
