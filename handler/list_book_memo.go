package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/naoyakurokawa/book_line_api/entity"
)

type ListBookMemo struct {
	Service   ListBookMemosService
	Validator *validator.Validate
}

type bookMemo struct {
	ID     entity.BookMemoID `json:"id"`
	BookID entity.BookID     `json:"book_id"`
	Page   int64             `json:"page"`
	Detail string            `json:"detail"`
}

func (lbm *ListBookMemo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body struct {
		BookID string `json:"book_id" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	err := lbm.Validator.Struct(body)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	book_memos, err := lbm.Service.ListBookMemos(ctx, body.BookID)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []bookMemo{}
	for _, bm := range book_memos {
		rsp = append(rsp, bookMemo{
			ID:     bm.ID,
			BookID: bm.BookID,
			Page:   bm.Page,
			Detail: bm.Detail,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
