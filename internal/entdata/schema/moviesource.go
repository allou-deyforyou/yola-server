package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MovieSource holds the schema definition for the MovieSource entity.
type MovieSource struct {
	ent.Schema
}

// Fields of the MovieSource.
func (MovieSource) Fields() []ent.Field {
	return []ent.Field{
		field.String("manga_serie_search_url").
			Optional().Nillable(),
		field.String("manga_film_search_url").
			Optional().Nillable(),
		field.String("serie_search_url").
			Optional().Nillable(),
		field.String("film_search_url").
			Optional().Nillable(),

		field.String("manga_serie_latest_url").
			Optional().Nillable(),
		field.String("manga_film_latest_url").
			Optional().Nillable(),
		field.String("serie_latest_url").
			Optional().Nillable(),
		field.String("film_latest_url").
			Optional().Nillable(),

		field.JSON("manga_serie_latest_post_selector", &MoviePostSelector{}).
			Optional(),
		field.JSON("manga_film_latest_post_selector", &MoviePostSelector{}).
			Optional(),
		field.JSON("serie_latest_post_selector", &MoviePostSelector{}).
			Optional(),
		field.JSON("film_latest_post_selector", &MoviePostSelector{}).
			Optional(),

		field.JSON("manga_serie_search_post_selector", &MoviePostSelector{}).
			Optional(),
		field.JSON("manga_film_search_post_selector", &MoviePostSelector{}).
			Optional(),
		field.JSON("serie_search_post_selector", &MoviePostSelector{}).
			Optional(),
		field.JSON("film_search_post_selector", &MoviePostSelector{}).
			Optional(),

		field.JSON("manga_serie_article_selector", &MovieArticleSelector{}).
			Optional(),
		field.JSON("manga_film_article_selector", &MovieArticleSelector{}).
			Optional(),
		field.JSON("serie_article_selector", &MovieArticleSelector{}).
			Optional(),
		field.JSON("film_article_selector", &MovieArticleSelector{}).
			Optional(),

		field.String("language").Default("fr"),
		field.Bool("status").Default(true),
		field.String("name").Unique(),
		field.String("url"),
	}
}

// Edges of the MovieSource.
func (MovieSource) Edges() []ent.Edge {
	return nil
}
