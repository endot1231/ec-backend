package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Admins struct {
	ent.Schema
}

// Fields of the User.
func (Admins) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
		field.Time("email_verified").Default(time.Now()).Optional().Nillable(),
		field.String("password"),
		field.String("remember_token").Optional(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.Time("deleted_at").Default(time.Now()).Optional().Nillable(),
	}
}

// Edges of the User.
func (Admins) Edges() []ent.Edge {
	return []ent.Edge{}
}
