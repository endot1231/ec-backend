package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type ProductBrands struct {
	ent.Schema
}

// Fields of the User.
func (ProductBrands) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.Time("deleted_at").Default(time.Now()).Optional().Nillable(),
	}
}

// Edges of the User.
func (ProductBrands) Edges() []ent.Edge {
	return []ent.Edge{}
}
