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

func (h *Handler) FilmLatestPost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	movieSources, _ := h.MovieSource.Query().Where(moviesource.Status(true)).All(ctx)
	sources := source.ParseMovieSources[source.FilmSource](movieSources)

	params := internal.Params(r.Form)
	page, _ := params.Int("page")

	var moviePosts []schema.MoviePost
	group := new(sync.WaitGroup)
	for _, s := range sources {
		group.Add(1)
		go func(source source.FilmSource) {
			posts := source.FilmLatestPost(ctx, page)
			moviePosts = append(moviePosts, posts...)
			group.Done()
		}(s)
	}
	group.Wait()

	response := internal.Shuffle(moviePosts)
	json.NewEncoder(w).Encode(response)
}
