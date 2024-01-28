package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Products struct {
	ent.Schema
}

// Fields of the User.
func (Products) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("shop_id"),
		field.Int64("product_category_id"),
		field.Int64("product_brand_id"),
		field.String("name"),
		field.String("description"),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.Time("deleted_at").Default(time.Now()).Optional().Nillable(),
	}
}

// Edges of the User.
func (Products) Edges() []ent.Edge {
	return []ent.Edge{}
}
