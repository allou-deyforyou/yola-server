package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"yola/handler"
	"yola/internal"
	"yola/internal/entdata"

	_ "github.com/mattn/go-sqlite3"
)

var httpHandler *handler.Handler
var entClient *entdata.Client
var port *string

func init() {
	entClient = internal.NewEntClient()
	httpHandler = handler.NewHandler(entClient)
	flag.StringVar(port, "PORT", "4000", "")
	flag.Parse()
}

func init() {
	httpHandler.Handle("/film/latest", handler.ParseHandler(httpHandler.FilmLatestPost))
}

func main() {
	defer entClient.Close()
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", *port), httpHandler))
}
