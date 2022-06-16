package source

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"yola/internal/element"
	"yola/internal/entdata"
	"yola/internal/entdata/schema"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

func chromeRequest(url string) (io.Reader, error) {
	opts := []chromedp.ExecAllocatorOption{
		// chromedp.Headless,
		// chromedp.DisableGPU,
		chromedp.NoSandbox,
	}
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()
	var response string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			res, er := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			response = res
			return er
		}),
	)
	return strings.NewReader(response), err
}

type VostfreeTvSource struct {
	*entdata.MovieSource
	*http.Client
}

func NewVostfreeTvSource(source *entdata.MovieSource) *VostfreeTvSource {
	return &VostfreeTvSource{
		Client:      http.DefaultClient,
		MovieSource: source,
	}
}

func (is *VostfreeTvSource) MangaLatestPost(_ context.Context, page int) []schema.MoviePost {
	response, err := chromeRequest(fmt.Sprintf("%s%s", is.URL, fmt.Sprintf(*is.MangaSerieLatestURL, page)))
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		return nil
	}
	return is.mangaLatestPostList(element.NewElement(document.Selection))
}

func (is *VostfreeTvSource) mangaLatestPostList(document *element.Element) []schema.MoviePost {
	selector := is.MangaSerieLatestPostSelector
	mangaList := make([]schema.MoviePost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *element.Element) {
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			link := element.ChildAttribute(selector.Link[0], selector.Link[1])
			title := element.ChildText(selector.Title[0])
			if strings.Contains(image, "imgur") {
				image = strings.ReplaceAll(image, path.Ext(image), "h"+path.Ext(image))
			}
			mangaList = append(mangaList, schema.MoviePost{
				Category: schema.MovieManga,
				Source:   is.Name,
				Image:    image,
				Title:    title,
				Link:     link,
			})
		})
	return mangaList
}

func (is *VostfreeTvSource) MangaArticle(_ context.Context, link string) *schema.MovieArticle {
	response, err := chromeRequest(link)
	if err != nil {
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		return nil
	}
	return is.mangaArticle(element.NewElement(document.Selection))
}

func (is *VostfreeTvSource) mangaArticle(document *element.Element) *schema.MovieArticle {
	articleSelector := is.MangaSerieArticleSelector

	description := document.ChildText(articleSelector.Description[0])
	genders := document.ChildTexts(articleSelector.Genders[0])

	var date string
	document.ForEachWithBreak(articleSelector.Date[0],
		func(i int, e *element.Element) bool {
			if strings.Contains(strings.ToLower(e.ChildText("span")), "année") {
				date = strings.TrimSpace(e.ChildText("a"))
				return false
			}
			return true
		})

	videos := make([]schema.MovieVideo, 0)
	document.ForEach(articleSelector.Hosters[0],
		func(i int, e *element.Element) {
			subtitleHosters := make([]string, 0)
			hosters := make([]string, 0)
			e.ForEach("div", func(i int, e *element.Element) {
				name := strings.ToLower(e.Text())
				link := document.ChildText(fmt.Sprintf("content_%v", e.Attribute("id")))
				switch name {
				case "sibnet":
					link = fmt.Sprintf("https://video.sibnet.ru/shell.php?videoid=%v", link)
				case "uqload":
					link = fmt.Sprintf("https://uqload.com/embed-%v.html", link)
				case "mytv":
					link = fmt.Sprintf("https://www.myvi.tv/embed/%v", link)
				}
				if strings.Contains(strings.ToLower(document.ChildText(articleSelector.Hosters[1])), "vostfr") {
					subtitleHosters = append(subtitleHosters, link)
				} else {
					hosters = append(hosters, link)
				}
			})
			videos = append(videos, schema.MovieVideo{
				Name:            strconv.Itoa(i + 1),
				SubtitleHosters: subtitleHosters,
				Hosters:         hosters,
			})
		})

	return &schema.MovieArticle{
		Sections: []schema.MovieSectionVideo{
			{Name: "Episodes", Videos: videos},
		},
		Description: description,
		Genders:     genders,
		Imdb:        "N/A",
		Date:        date,
	}
}