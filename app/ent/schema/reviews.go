package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Reviews struct {
	ent.Schema
}

// Fields of the User.
func (Reviews) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("contents"),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.Time("deleted_at").Default(time.Now()).Optional().Nillable(),
	}
}

// Edges of the User.
func (Reviews) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Users.Type).
			Ref("reviews"),
	}
}
