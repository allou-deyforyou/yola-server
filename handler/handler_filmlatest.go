package handler

import (
	"fmt"
	"net/http"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func (h *Handler) FilmLatestPost(w http.ResponseWriter, r *http.Request) {
	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).NoSandbox(true).MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage("https://vostfree.tv")
	fmt.Println(page.MustHTML())
	// ctx, cancel := context.WithTimeout(
	// 	context.Background(),
	// 	10*time.Second,
	// )
	// defer cancel()

	// params := internal.Params(r.Form)
	// page, _ := params.Int("page")

	// movieSources, _ := h.MovieSource.Query().Where(moviesource.Status(true)).All(ctx)
	// sources := source.ParseMovieSources[source.FilmSource](movieSources)

	// var moviePosts []schema.MoviePost
	// group := new(sync.WaitGroup)
	// for _, s := range sources {
	// 	group.Add(1)
	// 	go func(source source.FilmSource) {
	// 		posts := source.FilmLatestPost(ctx, page)
	// 		moviePosts = append(moviePosts, posts...)
	// 		group.Done()
	// 	}(s)
	// }
	// group.Wait()

	// response := internal.Shuffle(moviePosts)
	// json.NewEncoder(w).Encode(response)
}
