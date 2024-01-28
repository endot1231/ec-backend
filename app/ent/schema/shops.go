package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Shops holds the schema definition for the User entity.
type Shops struct {
	ent.Schema
}

// Fields of the Shops.
func (Shops) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Nillable(),
		field.String("address").Nillable(),
		field.String("email").Unique().NotEmpty(),
		field.Time("email_verified").Default(time.Now()).Optional().Nillable(),
		field.String("password").Optional().Sensitive(),
		field.String("remember_token").Optional(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.Time("deleted_at").Default(time.Now()).Optional().Nillable(),
	}
}

// Edges of the Shops.
func (Shops) Edges() []ent.Edge {
	return []ent.Edge{}
}
