package service

import (
	"context"
	"fmt"

	"github.com/naoyakurokawa/book_line_api/entity"
	"github.com/naoyakurokawa/book_line_api/store"
)

type FetchBookMemosService struct {
	DB   store.Queryer
	Repo BookMemoFetcher
}

func (fbm *FetchBookMemosService) FetchBookMemos(ctx context.Context, bookID string) (entity.BookMemos, error) {
	bms, err := fbm.Repo.FetchBookMemos(ctx, fbm.DB, bookID)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return bms, nil
}
