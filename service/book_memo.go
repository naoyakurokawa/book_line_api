package service

import (
	"context"
	"fmt"

	"github.com/naoyakurokawa/book_line_api/entity"
	"github.com/naoyakurokawa/book_line_api/store"
)

type CreateBookMemoService struct {
	DB   store.Execer
	Repo BookMemoCreator
}

func (cbm *CreateBookMemoService) CreateBookMemo(ctx context.Context, book_ID entity.BookID, page int64, detail string) error {
	bm := &entity.BookMemo{
		BookID: book_ID,
		Page:   page,
		Detail: detail,
	}
	if err := cbm.Repo.CreateBookMemo(ctx, cbm.DB, bm); err != nil {
		return fmt.Errorf("failed to create book memo: %w", err)
	}
	return nil
}

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
