package source

import (
	"context"
	"errors"

	"yola/internal/entdata"
	"yola/internal/entdata/schema"
)

type MangaSource interface {
	MangaArticle(link string) *schema.MovieArticle
	MangaLatestPostList(page int) []schema.MoviePost
	MangaSearchPostList(query string, page int) []schema.MoviePost
}

type SerieSource interface {
	SerieArticle(link string) *schema.MovieArticle
	SerieLatestPostList(page int) []schema.MoviePost
	SerieSearchPostList(query string, page int) []schema.MoviePost
}

type FilmSource interface {
	FilmArticle(link string) *schema.MovieArticle
	FilmLatestPost(ctx context.Context, page int) []schema.MoviePost
	FilmSearchPostList(query string, page int) []schema.MoviePost
}

func ParseMovieSources[T any](movieSources []*entdata.MovieSource) []T {
	var sources []T
	for _, movieSource := range movieSources {
		switch (interface{})(sources).(type) {
		case []MangaSource:
			if source, err := parseMangaSource(movieSource.Name, movieSource); err == nil {
				sources = append(sources, source.(T))
			}
		case []SerieSource:
			if source, err := parseSerieSource(movieSource.Name, movieSource); err == nil {
				sources = append(sources, source.(T))
			}
		case []FilmSource:
			if source, err := parseFilmSource(movieSource.Name, movieSource); err == nil {
				sources = append(sources, source.(T))
			}
		}
	}
	return sources
}

func parseMangaSource(name string, source *entdata.MovieSource) (MangaSource, error) {
	switch name {
	case "french-manga-net":
		return NewFrenchMangaNetSource(source), nil
	default:
		return nil, errors.New("no-found")
	}
}

func parseSerieSource(name string, source *entdata.MovieSource) (SerieSource, error) {
	switch name {
	case "french-stream-re":
		return NewFrenchStreamReSource(source), nil
	default:
		return nil, errors.New("no-found")
	}
}

func parseFilmSource(name string, source *entdata.MovieSource) (FilmSource, error) {
	switch name {
	case "french-stream-re":
		return NewFrenchStreamReSource(source), nil
	case "illimitestreaming-co":
		return NewIllimitestreamingco(source), nil
	default:
		return nil, errors.New("no-found")
	}
}
