package internal_test

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
	client, err := entdata.Open(dialect.SQLite, "../assets/files/yola.db?mode=memory&cache=shared&_fk=1")
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
		SetFilmLatestURL("/films/page/%v").
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

func TestCreateVostfreeTvSource(t *testing.T) {
	entClient.MovieSource.Create().
		SetStatus(true).
		SetName("vostfree-tv").
		SetURL("https://vostfree.tv/").
		SetMangaSerieLatestURL("/animes-vostfr/page/%v").
		SetMangaSerieLatestPostSelector(&schema.MoviePostSelector{
			Title: []string{".title"},
			Image: []string{"img", "src"},
			Link:  []string{"a", "href"},
			List:  []string{"#dle-content > .movie-poster"},
		}).
		SetMangaSerieSearchURL("/index.php?do=search").
		SetMangaSerieSearchPostSelector(&schema.MoviePostSelector{
			Title: []string{".title"},
			Image: []string{"img", "src"},
			Link:  []string{".title a", "href"},
			List:  []string{"#dle-content .search-result"},
		}).
		SetMangaSerieArticleSelector(&schema.MovieArticleSelector{
			Hosters:     []string{".new_player_selector option", ".new_player_bottom"},
			Genders:     []string{".slide-top li.right a"},
			Date:        []string{".slide-info p"},
			Description: []string{".slide-desc"},
		}).Save(context.Background())
}

func TestCreateRti1TV(t *testing.T) {
	entClient.Tv.Create().
		SetStatus(true).
		SetTitle("RTI 1").
		SetVideo("https://www.enovativecdn.com/rticdn/smil:rti1.smil/playlist.m3u8").
		SetLogo("https://rti.ci/_nuxt/img/6a12f31.svg").
		SetDescription("RTI 1 est la première chaîne de télévision généraliste publique ivoirienne qui émet en continu depuis Abidjan.").
		SaveX(context.Background())
}

func TestCreateRti2TV(t *testing.T) {
	entClient.Tv.Create().
		SetStatus(true).
		SetTitle("RTI 2").
		SetVideo("https://www.enovativecdn.com/rticdn/smil:rti2.smil/playlist.m3u8").
		SetLogo("https://rti.ci/_nuxt/img/6d85e57.svg").
		SetDescription("RTI 2 est une chaîne de télévision généraliste publique ivoirienne. Elle est consacrée à la jeunesse.").
		SaveX(context.Background())

}

func TestCreateRti3TV(t *testing.T) {
	entClient.Tv.Create().
		SetStatus(true).
		SetTitle("La 3").
		SetVideo("https://www.enovativecdn.com/rticdn/smil:rti3.smil/playlist.m3u8").
		SetLogo("https://rti.ci/_nuxt/img/4da62df.svg").
		SetDescription("RTI 3, baptisée La 3, est une chaîne de télévision ivoirienne du Groupe RTI lancée le 16 février 2020.").
		SaveX(context.Background())

}

func TestCreateNciTV(t *testing.T) {
	entClient.Tv.Create().
		SetStatus(true).
		SetTitle("NCI").
		SetVideo("https://nci-live.secure2.footprint.net/nci/nci.isml/.m3u8").
		SetLogo("https://static.wixstatic.com/media/f8668c_8cf416367fb743378ec26c7e7978a318~mv2_d_1692_1295_s_2.png").
		SetDescription("La Nouvelle Chaîne Ivoirienne, plus connue sous le sigle NCI est la première chaîne de télévision privée ivoirienne.").
		SaveX(context.Background())
}

func TestGetSources(t *testing.T) {
	log.Println(entClient.MovieSource.Query().AllX(context.Background()))
}

func TestRemoveSources(t *testing.T) {
	log.Println(entClient.MovieSource.Delete().Where(moviesource.Name("vostfree-tv")).Exec(context.Background()))
}
