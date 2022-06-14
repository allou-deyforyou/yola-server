package handler

import (
	"net/http"

	"yola/internal/entdata"
)

type Handler struct {
	*entdata.Client
	*http.ServeMux
}

func NewHandler(client *entdata.Client) *Handler {
	return &Handler{
		ServeMux: http.NewServeMux(),
		Client:   client,
	}
}

func ParseHandler(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		h.ServeHTTP(w, r)
	})
}
