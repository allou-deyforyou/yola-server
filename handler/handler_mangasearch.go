package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"yola/internal"
	"yola/internal/entdata/moviesource"
	"yola/internal/entdata/schema"
	"yola/internal/source"
)

func (h *Handler) MangaSearchPost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	params := internal.Params(r.Form)
	query, _ := params.String("query")
	page, _ := params.Int("page")

	movieSources := h.MovieSource.Query().Where(moviesource.Status(true)).AllX(ctx)
	sources := source.ParseMovieSources[source.MangaSource](movieSources)

	var moviePosts []schema.MoviePost
	group := new(sync.WaitGroup)
	for _, s := range sources {
		group.Add(1)
		go func(source source.MangaSource) {
			posts := source.MangaSearchPost(ctx, query, page)
			moviePosts = append(moviePosts, posts...)
			group.Done()
		}(s)
	}
	group.Wait()

	response := internal.Shuffle(moviePosts)
	json.NewEncoder(w).Encode(response)
}
