package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"yola/internal"
	"yola/internal/entdata/tv"
)

func (h *Handler) TvList(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	params := internal.Params(r.Form)
	language, _ := params.String("language")
	country, _ := params.String("country")

	tvQuery := h.Tv.Query()
	if len(language) != 0 {
		tvQuery = tvQuery.Where(tv.Language(language))
	}
	if len(country) != 0 {
		tvQuery = tvQuery.Where(tv.Country(country))
	}
	response := tvQuery.Where(tv.Status(true)).AllX(ctx)
	json.NewEncoder(w).Encode(response)
}
