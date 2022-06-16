package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"yola/internal"
	"yola/internal/entdata/moviesource"
	"yola/internal/source"
)

func (h *Handler) FilmArticle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	params := internal.Params(r.Form)
	name, _ := params.String("source")
	link, _ := params.String("link")

	movieSource := h.MovieSource.Query().Where(moviesource.And(moviesource.Status(true), moviesource.Name(name))).OnlyX(ctx)
	source, _ := source.ParseFilmSource(movieSource.Name, movieSource)

	response := source.FilmArticle(ctx, link)
	json.NewEncoder(w).Encode(response)
}
