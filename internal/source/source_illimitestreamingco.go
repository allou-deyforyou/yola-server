package source

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"yola/internal/element"
	"yola/internal/entdata"
	"yola/internal/entdata/schema"

	"github.com/PuerkitoBio/goquery"
)

type Illimitestreamingco struct {
	*entdata.MovieSource
	*http.Client
}

func NewIllimitestreamingco(source *entdata.MovieSource) *Illimitestreamingco {
	return &Illimitestreamingco{
		Client:      http.DefaultClient,
		MovieSource: source,
	}
}

func (is *Illimitestreamingco) FilmLatestPost(ctx context.Context, page int) []schema.MoviePost {
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(*is.FilmLatestURL, page)), nil)
	response, err := is.Do(request)
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmLatestPostList(element.NewElement(document.Selection))
}

func (is *Illimitestreamingco) filmLatestPostList(document *element.Element) []schema.MoviePost {
	selector := is.FilmLatestPostSelector
	filmList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *element.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])

			image = parseImage(image)
			filmList = append(filmList, schema.MoviePost{
				Category: schema.MovieFilm,
				Source:   is.Name,
				Image:    image,
				Title:    title,
				Link:     link,
			})
		})
	return filmList
}

func (is *Illimitestreamingco) FilmSearchPost(ctx context.Context, query string, page int) []schema.MoviePost {
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(*is.FilmSearchURL, page, query)), nil)
	response, err := is.Do(request)
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmSearchPost(element.NewElement(document.Selection))
}

func (is *Illimitestreamingco) filmSearchPost(document *element.Element) []schema.MoviePost {
	if !strings.Contains(strings.ToLower(document.ChildText(".movies-list-wrap .ml-title")), "recherche") {
		return nil
	}
	selector := is.FilmSearchPostSelector
	filmList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *element.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			image = parseImage(image)
			if strings.Contains(strings.ToLower(element.ChildText(".jtip-bottom")), "film") {
				filmList = append(filmList, schema.MoviePost{
					Category: schema.MovieFilm,
					Source:   is.Name,
					Image:    image,
					Title:    title,
					Link:     link,
				})
			}
		})
	return filmList
}

func (is *Illimitestreamingco) FilmArticle(ctx context.Context, link string) *schema.MovieArticle {
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	response, err := is.Do(request)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmArticle(element.NewElement(document.Selection))
}

func (is *Illimitestreamingco) filmArticle(document *element.Element) *schema.MovieArticle {
	articleSelector := is.FilmArticleSelector
	description := document.ChildText(articleSelector.Description[0])
	genders := make([]string, 0)
	document.ForEachWithBreak(articleSelector.Genders[0],
		func(i int, e *element.Element) bool {
			if strings.Contains(strings.ToLower(e.ChildText("strong")), "genre") {
				genders = append(genders, e.ChildTexts("a")...)
				return false
			}
			return true
		})

	var date string
	document.ForEachWithBreak(articleSelector.Date[0],
		func(i int, e *element.Element) bool {
			if strings.Contains(strings.ToLower(e.ChildText("strong")), "année") {
				date = strings.TrimSpace(e.ChildText("a"))
				return false
			}
			return true
		})

	var imdb string
	document.ForEachWithBreak(articleSelector.Imdb[0],
		func(i int, e *element.Element) bool {
			if strings.Contains(strings.ToLower(e.ChildText("strong")), "imdb") {
				imdb = strings.TrimSpace(e.ChildText("span"))
				return false
			}
			return true
		})

	videos := make([]schema.MovieVideo, 0)
	subtitleHosters := make([]string, 0)
	hosters := make([]string, 0)
	document.ForEach(articleSelector.Hosters[0],
		func(i int, e *element.Element) {
			id := e.ChildAttribute("a", "href")
			if id != "" {
				link := strings.ToLower(document.ChildText(id))
				if strings.TrimSpace(strings.ToLower(e.ChildText("span:nth-child(2) h6"))) == "vf" {
					hosters = append(hosters, link)
				} else {
					subtitleHosters = append(subtitleHosters, link)
				}
			}
		})
	videos = append(videos, schema.MovieVideo{
		SubtitleHosters: subtitleHosters,
		Hosters:         hosters,
		Name:            "Video",
	})

	if len(genders) == 0 {
		genders = append(genders, "N/A")
	}
	return &schema.MovieArticle{
		Description: description,
		Genders:     genders,
		Sections: []schema.MovieSectionVideo{
			{Name: "Film", Videos: videos},
		},
		Imdb: imdb,
		Date: date,
	}
}
