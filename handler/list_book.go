package handler

import (
	"net/http"

	"github.com/naoyakurokawa/book_line_api/entity"
)

type ListBook struct {
	Service ListBooksService
}

type book struct {
	ID    entity.BookID `json:"id"`
	Title string        `json:"title"`
	Isbn  entity.Isbn   `json:"isbn"`
}

func (lt *ListBook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	books, err := lt.Service.ListBooks(ctx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []book{}
	for _, b := range books {
		rsp = append(rsp, book{
			ID:    b.ID,
			Title: b.Title,
			Isbn:  b.Isbn,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
