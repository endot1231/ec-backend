package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type ProductSizes struct {
	ent.Schema
}

// Fields of the User.
func (ProductSizes) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),
		field.String("size"),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.Time("deleted_at").Default(time.Now()).Optional().Nillable(),
	}
}

// Edges of the User.
func (ProductSizes) Edges() []ent.Edge {
	return []ent.Edge{}
}
