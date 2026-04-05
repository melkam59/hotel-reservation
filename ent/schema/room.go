package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("size"),
		field.Bool("seaside"),
		field.Float("price"),
		field.Int("hotel_id"),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hotel", Hotel.Type).
			Ref("rooms").
			Unique().
			Field("hotel_id").
			Required(),
		edge.To("bookings", Booking.Type),
	}
}
