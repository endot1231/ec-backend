package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type ProductStock struct {
	ent.Schema
}

// Fields of the User.
func (ProductStock) Fields() []ent.Field {
	return []ent.Field{
		field.String("description"),
		field.Int("price").Positive(),
		field.Int("stock_quantity").Positive(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the User.
func (ProductStock) Edges() []ent.Edge {
	return []ent.Edge{}
}
