package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("room_id"),
		field.Int("num_persons"),
		field.Time("from_date"),
		field.Time("till_date"),
		field.Bool("canceled").Default(false),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("bookings").
			Unique().
			Field("user_id").
			Required(),
		edge.From("room", Room.Type).
			Ref("bookings").
			Unique().
			Field("room_id").
			Required(),
	}
}
