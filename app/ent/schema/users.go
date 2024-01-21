package schema

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	gen "github.com/endot1231/ec-backend/ent"
	"github.com/endot1231/ec-backend/ent/hook"

	"github.com/endot1231/ec-backend/secrets"
)

// User holds the schema definition for the User entity.
type Users struct {
	ent.Schema
}

// Fields of the User.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email").Unique(),
		field.Time("email_verified").Default(time.Now()).Optional().Nillable(),
		field.String("password").Optional().Sensitive(),
		field.String("remember_token").Optional(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.Time("deleted_at").Default(time.Now()).Optional().Nillable(),
	}
}

func (Users) Indexes() []ent.Index {
	return []ent.Index{}
}

// Edges of the User.
func (Users) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reviews", Reviews.Type),
	}
}

// Hooks of the User.
func (Users) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UsersFunc(func(ctx context.Context, m *gen.UsersMutation) (gen.Value, error) {
					v, ok := m.Password()
					if !ok || v == "" {
						return nil, errors.New("unexpected 'password' value")
					}
					c, err := secrets.Encrypt(v)
					if err != nil {
						return nil, err
					}

					m.SetPassword(c)
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
	}
}
