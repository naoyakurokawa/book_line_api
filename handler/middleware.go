package handler

import (
	"log"
	"net/http"

	"github.com/naoyakurokawa/book_line_api/auth"
	"github.com/naoyakurokawa/book_line_api/config"
	"github.com/rs/cors"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("AdminMiddleware")
		if !auth.IsAdmin(r.Context()) {
			RespondJSON(r.Context(), w, ErrResponse{
				Message: "not admin",
			}, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(j *auth.JWTer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req, err := j.FillContext(r)
			if err != nil {
				RespondJSON(r.Context(), w, ErrResponse{
					Message: "not find auth info",
					Details: []string{err.Error()},
				}, http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}

func CorsMiddlewareFunc() func(http.Handler) http.Handler {
	cfg, _ := config.New()
	corsWrapper := cors.New(cors.Options{
		AllowedOrigins:   cfg.Cors,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Accept-Language", "Authorization"},
		AllowCredentials: true,
	})

	return corsWrapper.Handler
}
