package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"yola/handler"
	"yola/internal"
	"yola/internal/entdata"

	_ "github.com/mattn/go-sqlite3"
)

var httpHandler *handler.Handler
var entClient *entdata.Client
var port string

func init() {
	entClient = internal.NewEntClient()
	httpHandler = handler.NewHandler(entClient)
	port = os.Getenv("PORT")
}

func init() {
	httpHandler.Handle("/film/article", handler.ParseHandler(httpHandler.FilmArticle))
	httpHandler.Handle("/film/latest", handler.ParseHandler(httpHandler.FilmLatestPost))
	httpHandler.Handle("/search/film", handler.ParseHandler(httpHandler.FilmSearchPost))

	httpHandler.Handle("/serie/article", handler.ParseHandler(httpHandler.SerieArticle))
	httpHandler.Handle("/serie/latest", handler.ParseHandler(httpHandler.SerieLatestPost))
	httpHandler.Handle("/search/serie", handler.ParseHandler(httpHandler.SerieSearchPost))

	httpHandler.Handle("/manga/article", handler.ParseHandler(httpHandler.MangaArticle))
	httpHandler.Handle("/manga/latest", handler.ParseHandler(httpHandler.MangaLatestPost))
	httpHandler.Handle("/search/manga", handler.ParseHandler(httpHandler.MangaSearchPost))
}

func main() {
	defer entClient.Close()
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), httpHandler))
}
