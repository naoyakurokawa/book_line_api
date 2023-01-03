package store

import (
	"context"

	"github.com/naoyakurokawa/book_line_api/entity"
)

func (r *Repository) FetchBookMemos(
	ctx context.Context, db Queryer, bookID string,
) (entity.BookMemos, error) {
	book_memos := entity.BookMemos{}
	sql := `SELECT 
				id, book_id, page, detail,
				created, modified 
			FROM book_memos WHERE book_id = ?`
	if err := db.SelectContext(ctx, &book_memos, sql, bookID); err != nil {
		return nil, err
	}
	return book_memos, nil
}

func (r *Repository) CreateBookMemo(
	ctx context.Context, db Execer, bm *entity.BookMemo,
) error {
	bm.Created = r.Clocker.Now()
	bm.Modified = r.Clocker.Now()
	sql := `INSERT INTO book_memos(book_id, page, detail, created, modified)
		VALUES (?, ?, ?, ?, ?) `
	_, err := db.ExecContext(ctx, sql, bm.BookID, bm.Page, bm.Detail, bm.Created, bm.Created)
	if err != nil {
		return err
	}
	return nil
}
