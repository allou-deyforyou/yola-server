package source

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"yola/internal/element"
	"yola/internal/entdata"
	"yola/internal/entdata/schema"

	"github.com/PuerkitoBio/goquery"
)

type FrenchStreamReSource struct {
	*entdata.MovieSource
	*http.Client
}

func NewFrenchStreamReSource(source *entdata.MovieSource) *FrenchStreamReSource {
	return &FrenchStreamReSource{
		Client:      http.DefaultClient,
		MovieSource: source,
	}
}

func (is *FrenchStreamReSource) FilmLatestPost(ctx context.Context, page int) []schema.MoviePost {
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(*is.FilmLatestURL, page)), nil)
	response, err := is.Do(request)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.filmLatestPost(element.NewElement(document.Selection))
}

func (is *FrenchStreamReSource) filmLatestPost(document *element.Element) []schema.MoviePost {
	selector := is.FilmLatestPostSelector
	filmList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *element.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			if strings.Contains(image, "imgur") {
				image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
			}
			if strings.Contains(image, "tmdb") {
				_, file := path.Split(image)
				image = fmt.Sprintf("https://image.tmdb.org/t/p/w500/%s", file)
			}
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

func (is *FrenchStreamReSource) FilmSearchPost(ctx context.Context, query string, page int) []schema.MoviePost {
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://french-manga.net/index.php?do=search",
		strings.NewReader(url.Values{
			"do":           []string{"search"},
			"subaction":    []string{"search"},
			"story":        []string{query},
			"search_start": []string{strconv.Itoa(page)},
			"full_search":  []string{"1"},
			"result_from":  []string{strconv.Itoa(page)},
			"titleonly":    []string{"3"},
			"replyless":    []string{"0"},
			"replylimit":   []string{"0"},
			"searchdate":   []string{"0"},
			"beforeafter":  []string{"after"},
			"sortby":       []string{"date"},
			"resorder":     []string{"desc"},
			"showposts":    []string{"0"},
			"catlist[]":    []string{"9"},
		}.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

func (is *FrenchStreamReSource) filmSearchPost(document *element.Element) []schema.MoviePost {
	selector := is.FilmSearchPostSelector
	filmList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *element.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			if strings.Contains(image, "imgur") {
				image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
			}
			if strings.Contains(image, "tmdb") {
				_, file := path.Split(image)
				image = fmt.Sprintf("https://image.tmdb.org/t/p/w500/%s", file)
			}
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

func (is *FrenchStreamReSource) SerieLatestPost(ctx context.Context, page int) []schema.MoviePost {
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(*is.SerieLatestURL, page)), nil)
	response, err := is.Do(request)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.serieLatestPost(element.NewElement(document.Selection))
}

func (is *FrenchStreamReSource) serieLatestPost(document *element.Element) []schema.MoviePost {
	selector := is.SerieLatestPostSelector
	serieList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *element.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			if strings.Contains(image, "imgur") {
				image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
			}
			if strings.Contains(image, "tmdb") {
				_, file := path.Split(image)
				image = fmt.Sprintf("https://image.tmdb.org/t/p/w500/%s", file)
			}
			serieList = append(serieList, schema.MoviePost{
				Category: schema.MovieSerie,
				Source:   is.Name,
				Image:    image,
				Title:    title,
				Link:     link,
			})
		})
	return serieList
}

func (is *FrenchStreamReSource) SerieSearchPost(ctx context.Context, query string, page int) []schema.MoviePost {
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://french-manga.net/index.php?do=search",
		strings.NewReader(url.Values{
			"do":           []string{"search"},
			"subaction":    []string{"search"},
			"story":        []string{query},
			"search_start": []string{strconv.Itoa(page)},
			"full_search":  []string{"1"},
			"result_from":  []string{strconv.Itoa(page)},
			"titleonly":    []string{"3"},
			"replyless":    []string{"0"},
			"replylimit":   []string{"0"},
			"searchdate":   []string{"0"},
			"beforeafter":  []string{"after"},
			"sortby":       []string{"date"},
			"resorder":     []string{"desc"},
			"showposts":    []string{"0"},
			"catlist[]":    []string{"10"},
		}.Encode()),
	)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := is.Do(request)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	return is.serieSearchPost(element.NewElement(document.Selection))
}

func (is *FrenchStreamReSource) serieSearchPost(document *element.Element) []schema.MoviePost {
	selector := is.SerieSearchPostSelector
	serieList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *element.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			if strings.Contains(image, "imgur") {
				image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
			}
			serieList = append(serieList, schema.MoviePost{
				Category: schema.MovieSerie,
				Source:   is.Name,
				Image:    image,
				Title:    title,
				Link:     link,
			})
		})
	return serieList
}

func (is *FrenchStreamReSource) FilmArticle(ctx context.Context, link string) *schema.MovieArticle {
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

func (is *FrenchStreamReSource) filmArticle(document *element.Element) *schema.MovieArticle {
	articleSelector := is.FilmArticleSelector

	description := document.ChildText(articleSelector.Description[0])
	// imdb := document.ChildText(articleSelector.Imdb[0])

	genders := make([]string, 0)
	document.ForEachWithBreak(articleSelector.Genders[0],
		func(i int, e *element.Element) bool {
			if strings.Contains(e.ChildText("span"), "Genre") {
				genders = append(genders, e.ChildTexts("a")...)
				return false
			}
			return true
		})

	var date string
	document.ForEachWithBreak(articleSelector.Date[0],
		func(i int, e *element.Element) bool {
			if strings.Contains(e.ChildText("span"), "sortie") {
				date = strings.TrimSpace(e.Selection.Contents().Not("span").Text())
				return false
			}
			return true
		})

	videos := make([]schema.MovieVideo, 0)

	subtitleHosters := make([]string, 0)
	hosters := make([]string, 0)
	document.ForEach(articleSelector.Hosters[0],
		func(i int, e *element.Element) {
			if strings.Contains(strings.ToLower(strings.TrimSpace(e.ChildText("li"))), "vostfr") {
				subtitleHosters = append(subtitleHosters, e.ChildAttribute("li a", "href"))
			}
			if strings.Contains(strings.ToLower(strings.TrimSpace(e.ChildText("li"))), "french") {
				hosters = append(hosters, e.ChildAttribute("li a", "href"))
			}
		})
	videos = append(videos, schema.MovieVideo{
		SubtitleHosters: subtitleHosters,
		Hosters:         hosters,
		Name:            "Film",
	})

	if len(genders) == 0 {
		genders = append(genders, "N/A")
	}
	return &schema.MovieArticle{
		Description: description,
		Genders:     genders,
		Sections: []schema.MovieSectionVideo{
			{Name: "Episodes", Videos: videos},
		},
		Imdb: "N/A",
		Date: date,
	}
}

func (is *FrenchStreamReSource) SerieArticle(ctx context.Context, link string) *schema.MovieArticle {
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
	return is.serieArticle(element.NewElement(document.Selection))
}

func (is *FrenchStreamReSource) serieArticle(document *element.Element) *schema.MovieArticle {
	articleSelector := is.SerieArticleSelector

	description := document.ChildText(articleSelector.Description[0])

	genders := make([]string, 0)
	document.ForEachWithBreak(articleSelector.Genders[0],
		func(i int, e *element.Element) bool {
			if strings.Contains(e.ChildText("span"), "Genre") {
				genders = append(genders, strings.Split(strings.TrimSpace(e.Selection.Contents().Not("span").Text()), ", ")...)
				return false
			}
			return true
		})

	videos := make([]schema.MovieVideo, 0)

	videosMap := make(map[string]schema.MovieVideo)
	document.ForEach(articleSelector.Hosters[0],
		func(index int, version *element.Element) {
			version.ForEach(articleSelector.Hosters[1], func(i int, episode *element.Element) {
				id := strings.TrimSpace(strings.TrimPrefix(strings.ToLower(episode.Attribute("title")), "episode"))
				video := schema.MovieVideo{Name: id, Hosters: make([]string, 0), SubtitleHosters: make([]string, 0)}
				if v, ok := videosMap[id]; ok {
					video = v
				}
				ref := episode.Attribute("data-rel")
				if index == 0 {
					if ref == "" {
						video.Hosters = append(video.Hosters, episode.Attribute("href"))
					} else {
						document.ForEach(fmt.Sprintf("#%v li a", ref), func(i int, hoster *element.Element) {
							link := hoster.Attribute("href")
							if link == "" {
								video.Hosters = append(video.Hosters, episode.Attribute("href"))
							} else {
								video.Hosters = append(video.Hosters, link)
							}
						})
					}
				} else {
					if ref == "" {
						video.SubtitleHosters = append(video.SubtitleHosters, episode.Attribute("href"))
					} else {
						document.ForEach(fmt.Sprintf("#%v li a", ref), func(i int, hoster *element.Element) {
							link := hoster.Attribute("href")
							if link == "" {
								video.Hosters = append(video.Hosters, episode.Attribute("href"))
							} else {
								video.SubtitleHosters = append(video.SubtitleHosters, link)
							}
						})
					}
				}
				videosMap[id] = video
			})
		})
	for _, v := range videosMap {
		videos = append(videos, v)
	}
	if len(genders) == 0 {
		genders = append(genders, "N/A")
	}
	return &schema.MovieArticle{
		Description: description,
		Genders:     genders,
		Sections: []schema.MovieSectionVideo{
			{Name: "Episodes", Videos: videos},
		},
		Imdb: "N/A",
		Date: "N/A",
	}
}
