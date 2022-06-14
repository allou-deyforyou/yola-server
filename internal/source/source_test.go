package source_test

import (
	"context"
	"log"
	"testing"

	"yola/internal/entdata"
	"yola/internal/entdata/migrate"
	"yola/internal/entdata/moviesource"
	"yola/internal/entdata/schema"

	"entgo.io/ent/dialect"

	_ "github.com/mattn/go-sqlite3"
)

var entClient *entdata.Client

func init() {
	client, err := entdata.Open(dialect.SQLite, "../../yola.db?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	ctx := context.Background()
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	entClient = client
}

func TestCreateIllimitestreamingcoSource(t *testing.T) {
	entClient.MovieSource.Create().
		SetStatus(true).
		SetName("illimitestreaming-co").
		SetURL("https://www.illimitestreaming.co").
		SetFilmLatestURL("/film/page/%v").
		SetFilmLatestPostSelector(&schema.MoviePostSelector{
			Title: []string{"h2"},
			Image: []string{"img", "data-original"},
			Link:  []string{"a", "href"},
			List:  []string{".movies-list .ml-item"},
		}).
		SetFilmSearchURL("/page/%v?s=%v").
		SetFilmSearchPostSelector(&schema.MoviePostSelector{
			Title: []string{"h2"},
			Image: []string{"img", "data-original"},
			Link:  []string{"a", "href"},
			List:  []string{".movies-list .ml-item"},
		}).
		SetFilmArticleSelector(&schema.MovieArticleSelector{
			Imdb:        []string{".mvic-desc p"},
			Genders:     []string{".mvic-desc p"},
			Date:        []string{".mvic-desc p"},
			Hosters:     []string{".idTabs > div"},
			Description: []string{".mvic-desc .desc"},
		}).Save(context.Background())
}

func TestCreateFrenchStreamSource(t *testing.T) {
	entClient.MovieSource.Create().
		SetStatus(true).
		SetName("french-stream-re").
		SetURL("https://french-stream.re").
		SetFilmLatestURL("/film/page/%v").
		SetFilmLatestPostSelector(&schema.MoviePostSelector{
			Title: []string{".short-poster .short-title"},
			Image: []string{".short-poster img", "src"},
			Link:  []string{"a.short-poster", "href"},
			List:  []string{"#dle-content > .short"},
		}).
		SetFilmSearchURL("/index.php?do=search").
		SetFilmSearchPostSelector(&schema.MoviePostSelector{
			Title: []string{".short-poster .short-title"},
			Image: []string{".short-poster img", "src"},
			Link:  []string{"a.short-poster", "href"},
			List:  []string{"#dle-content > .short"},
		}).
		SetFilmArticleSelector(&schema.MovieArticleSelector{
			Hosters:     []string{"#primary_nav_wrap > ul > li"},
			Imdb:        []string{".fmain .frate .fr-count"},
			Genders:     []string{".fmain .flist li"},
			Date:        []string{".fmain .flist li"},
			Description: []string{".fmain .fdesc"},
		}).
		SetSerieLatestURL("/serie/page/%v").
		SetSerieLatestPostSelector(&schema.MoviePostSelector{
			Title: []string{".short-poster .short-title"},
			Image: []string{".short-poster img", "src"},
			Link:  []string{"a.short-poster", "href"},
			List:  []string{"#dle-content > .short"},
		}).
		SetSerieSearchURL("/index.php?do=search").
		SetSerieSearchPostSelector(&schema.MoviePostSelector{
			Title: []string{".short-poster .short-title"},
			Image: []string{".short-poster img", "src"},
			Link:  []string{"a.short-poster", "href"},
			List:  []string{"#dle-content > .short"},
		}).
		SetSerieArticleSelector(&schema.MovieArticleSelector{
			Hosters:     []string{".elink", "a"},
			Imdb:        []string{".fmain .frate .fr-count"},
			Genders:     []string{".fmain .flist li"},
			Date:        []string{".fmain .flist li"},
			Description: []string{".fmain .fdesc"},
		}).Save(context.Background())
}

func TestCreateFrenchMangaNetSource(t *testing.T) {
	entClient.MovieSource.Create().
		SetStatus(true).
		SetName("french-manga-net").
		SetURL("https://french-manga.net").
		SetMangaSerieLatestURL("/manga-streaming/page/%v").
		SetMangaSerieLatestPostSelector(&schema.MoviePostSelector{
			Title: []string{".short-poster .short-title"},
			Image: []string{".short-poster img", "src"},
			Link:  []string{"a.short-poster", "href"},
			List:  []string{"#dle-content > .short"},
		}).
		SetMangaSerieSearchURL("/index.php?do=search").
		SetMangaSerieSearchPostSelector(&schema.MoviePostSelector{
			Title: []string{".short-poster .short-title"},
			Image: []string{".short-poster img", "src"},
			Link:  []string{"a.short-poster", "href"},
			List:  []string{"#dle-content > .short"},
		}).
		SetMangaSerieArticleSelector(&schema.MovieArticleSelector{
			Hosters:     []string{".elink", "a"},
			Imdb:        []string{".fmain .frate .fr-count"},
			Genders:     []string{".fmain .flist li"},
			Date:        []string{".fmain .flist li"},
			Description: []string{".fmain .fdesc"},
		}).Save(context.Background())
}

func TestGetSources(t *testing.T) {
	log.Println(entClient.MovieSource.Query().Where(moviesource.Status(true)).AllX(context.Background()))
}
