package handler

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/chromedp/chromedp"
)

func (h *Handler) FilmLatestPost(w http.ResponseWriter, r *http.Request) {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.Headless,
		chromedp.DisableGPU,
		chromedp.NoSandbox,
	}
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://uqload.com/embed-wbcmmtoipiyf.html`),
		chromedp.OuterHTML("body", &res),
	)
	if err != nil {
		log.Fatal("Allou Error: ", err)
	}

	log.Println("Allou Data: ", strings.TrimSpace(res))

	// ctx, cancel := context.WithTimeout(
	// 	context.Background(),
	// 	10*time.Second,
	// )
	// defer cancel()

	// movieSources, _ := h.MovieSource.Query().Where(moviesource.Status(true)).All(ctx)
	// sources := source.ParseMovieSources[source.FilmSource](movieSources)

	// params := internal.Params(r.Form)
	// page, _ := params.Int("page")
	// log.Println(page)

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
