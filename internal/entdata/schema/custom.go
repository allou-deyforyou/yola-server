package schema

type MovieCategory string

const (
	MovieFilm  MovieCategory = "film"
	MovieSerie MovieCategory = "serie"
	MovieManga MovieCategory = "manga"
)

type MoviePost struct {
	Category MovieCategory `json:"category"`
	Source   string        `json:"source"`
	Title    string        `json:"title"`
	Image    string        `json:"image"`
	Link     string        `json:"link"`
}

type MovieSectionVideo struct {
	Videos []MovieVideo `json:"videos"`
	Name   string       `json:"name"`
}

type MovieVideo struct {
	SubtitleHosters []string `json:"subtitle_hosters"`
	Hosters         []string `json:"hosters"`
	Name            string   `json:"name"`
}

type MovieArticle struct {
	Description string              `json:"description"`
	Sections    []MovieSectionVideo `json:"sections"`
	Genders     []string            `json:"genders"`
	Date        string              `json:"date"`
	Imdb        string              `json:"imdb"`
}

type MovieArticleSelector struct {
	Description []string `json:"description,omitempty"`
	Genders     []string `json:"genders,omitempty"`
	Hosters     []string `json:"hosters,omitempty"`
	Date        []string `json:"date,omitempty"`
	Imdb        []string `json:"imdb,omitempty"`
}

type MoviePostSelector struct {
	Title []string `json:"title,omitempty"`
	Image []string `json:"image,omitempty"`
	List  []string `json:"list,omitempty"`
	Link  []string `json:"link,omitempty"`
}
