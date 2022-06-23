package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"yola/handler"
	"yola/internal"

	_ "github.com/mattn/go-sqlite3"
)

var server *handler.Handler

func init() {
	server = handler.NewHandler(internal.NewEntClient())
}

func init() {
	server.Handle("/tv/list", handler.ParseHandler(server.TvList))

	server.Handle("/film/article", handler.ParseHandler(server.FilmArticle))
	server.Handle("/film/latest", handler.ParseHandler(server.FilmLatestPost))
	server.Handle("/search/film", handler.ParseHandler(server.FilmSearchPost))

	server.Handle("/serie/article", handler.ParseHandler(server.SerieArticle))
	server.Handle("/serie/latest", handler.ParseHandler(server.SerieLatestPost))
	server.Handle("/search/serie", handler.ParseHandler(server.SerieSearchPost))

	server.Handle("/manga/article", handler.ParseHandler(server.MangaArticle))
	server.Handle("/manga/latest", handler.ParseHandler(server.MangaLatestPost))
	server.Handle("/search/manga", handler.ParseHandler(server.MangaSearchPost))
}

func main() {
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), server))
}
