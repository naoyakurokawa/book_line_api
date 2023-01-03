package handler

import (
	"net/http"

	"github.com/naoyakurokawa/book_line_api/entity"
)

type FetchBooksHandler struct {
	Service FetchBooksService
}

type book struct {
	ID   entity.BookID `json:"id"`
	Isbn entity.Isbn   `json:"isbn"`
}

func (fb *FetchBooksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	books, err := fb.Service.FetchBooks(ctx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []book{}
	for _, b := range books {
		rsp = append(rsp, book{
			ID:   b.ID,
			Isbn: b.Isbn,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
