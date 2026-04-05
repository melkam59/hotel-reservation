package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Hotel holds the schema definition for the Hotel entity.
type Hotel struct {
	ent.Schema
}

// Fields of the Hotel.
func (Hotel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("location"),
		field.Int("rating"),
	}
}

// Edges of the Hotel.
func (Hotel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rooms", Room.Type),
	}
}
