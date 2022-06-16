package source

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"path"
	"strings"

	"yola/internal/entdata"
	"yola/internal/entdata/schema"
)

type MangaSource interface {
	MangaArticle(ctx context.Context, link string) *schema.MovieArticle
	MangaLatestPost(ctx context.Context, page int) []schema.MoviePost
	MangaSearchPost(ctx context.Context, query string, page int) []schema.MoviePost
}

type SerieSource interface {
	SerieArticle(ctx context.Context, link string) *schema.MovieArticle
	SerieLatestPost(ctx context.Context, page int) []schema.MoviePost
	SerieSearchPost(ctx context.Context, query string, page int) []schema.MoviePost
}

type FilmSource interface {
	FilmArticle(ctx context.Context, link string) *schema.MovieArticle
	FilmLatestPost(ctx context.Context, page int) []schema.MoviePost
	FilmSearchPost(ctx context.Context, query string, page int) []schema.MoviePost
}

func ParseMovieSources[T any](movieSources []*entdata.MovieSource) []T {
	var sources []T
	for _, movieSource := range movieSources {
		switch (interface{})(sources).(type) {
		case []MangaSource:
			if source, err := ParseMangaSource(movieSource.Name, movieSource); err == nil {
				sources = append(sources, source.(T))
			}
		case []SerieSource:
			if source, err := ParseSerieSource(movieSource.Name, movieSource); err == nil {
				sources = append(sources, source.(T))
			}
		case []FilmSource:
			if source, err := ParseFilmSource(movieSource.Name, movieSource); err == nil {
				sources = append(sources, source.(T))
			}
		}
	}
	return sources
}

func ParseMangaSource(name string, source *entdata.MovieSource) (MangaSource, error) {
	switch name {
	case "french-manga-net":
		return NewFrenchMangaNetSource(source), nil
	case "vostfree-tv":
		return NewVostfreeTvSource(source), nil
	default:
		return nil, errors.New("no-found")
	}
}

func ParseSerieSource(name string, source *entdata.MovieSource) (SerieSource, error) {
	switch name {
	case "french-stream-re":
		return NewFrenchStreamReSource(source), nil
	default:
		return nil, errors.New("no-found")
	}
}

func ParseFilmSource(name string, source *entdata.MovieSource) (FilmSource, error) {
	switch name {
	case "french-stream-re":
		return NewFrenchStreamReSource(source), nil
	case "illimitestreaming-co":
		return NewIllimitestreamingco(source), nil
	default:
		return nil, errors.New("no-found")
	}
}

func parseImage(image string) string {
	if strings.Contains(image, "imgur") {
		image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
	}
	if strings.Contains(image, "tmdb") {
		_, file := path.Split(image)
		image = fmt.Sprintf("https://image.tmdb.org/t/p/w500/%s", file)
	}
	return image
}

func parseURL(baseURL, rawURL string) string {
	bu, _ := url.Parse(baseURL)
	u, err := bu.Parse(rawURL)
	if err != nil {
		log.Println(err)
		return rawURL
	}
	return u.String()
}
