package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/naoyakurokawa/book_line_api/entity"
)

type FetchBooksHandler struct {
	Service   FetchBooksServicer
	Validator *validator.Validate
}

type book struct {
	ID   entity.BookID `json:"id"`
	Isbn int64         `json:"isbn"`
}

func (fb *FetchBooksHandler) FetchBooks(w http.ResponseWriter, r *http.Request) {
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

type CreateBooksHandler struct {
	Service   CreateBookServicer
	Validator *validator.Validate
}

func (cb *CreateBooksHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body struct {
		Isbn string `json:"isbn" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := cb.Validator.Struct(body); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	isbn, err := strconv.ParseInt(body.Isbn, 10, 64)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	err = cb.Service.CreateBook(ctx, isbn)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	rsp := struct {
		message string
	}{message: "success create book"}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
