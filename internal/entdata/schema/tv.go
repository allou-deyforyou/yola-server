package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tv holds the schema definition for the Tv entity.
type Tv struct {
	ent.Schema
}

// Fields of the Tv.
func (Tv) Fields() []ent.Field {
	return []ent.Field{
		field.String("logo").NotEmpty(),
		field.String("video").NotEmpty(),
		field.String("title").NotEmpty(),
		field.Bool("status").Default(true),
		field.String("country").Default("ci"),
		field.String("description").NotEmpty(),
		field.String("language").Default("fr"),
	}
}

// Edges of the Tv.
func (Tv) Edges() []ent.Edge {
	return nil
}
