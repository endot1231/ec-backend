// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/endot1231/ec-backend/ent/shops"
)

// ShopsCreate is the builder for creating a Shops entity.
type ShopsCreate struct {
	config
	mutation *ShopsMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *ShopsCreate) SetName(s string) *ShopsCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetAddress sets the "address" field.
func (sc *ShopsCreate) SetAddress(s string) *ShopsCreate {
	sc.mutation.SetAddress(s)
	return sc
}

// SetEmail sets the "email" field.
func (sc *ShopsCreate) SetEmail(s string) *ShopsCreate {
	sc.mutation.SetEmail(s)
	return sc
}

// SetEmailVerified sets the "email_verified" field.
func (sc *ShopsCreate) SetEmailVerified(t time.Time) *ShopsCreate {
	sc.mutation.SetEmailVerified(t)
	return sc
}

// SetNillableEmailVerified sets the "email_verified" field if the given value is not nil.
func (sc *ShopsCreate) SetNillableEmailVerified(t *time.Time) *ShopsCreate {
	if t != nil {
		sc.SetEmailVerified(*t)
	}
	return sc
}

// SetPassword sets the "password" field.
func (sc *ShopsCreate) SetPassword(s string) *ShopsCreate {
	sc.mutation.SetPassword(s)
	return sc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (sc *ShopsCreate) SetNillablePassword(s *string) *ShopsCreate {
	if s != nil {
		sc.SetPassword(*s)
	}
	return sc
}

// SetRememberToken sets the "remember_token" field.
func (sc *ShopsCreate) SetRememberToken(s string) *ShopsCreate {
	sc.mutation.SetRememberToken(s)
	return sc
}

// SetNillableRememberToken sets the "remember_token" field if the given value is not nil.
func (sc *ShopsCreate) SetNillableRememberToken(s *string) *ShopsCreate {
	if s != nil {
		sc.SetRememberToken(*s)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *ShopsCreate) SetCreatedAt(t time.Time) *ShopsCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *ShopsCreate) SetNillableCreatedAt(t *time.Time) *ShopsCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *ShopsCreate) SetUpdatedAt(t time.Time) *ShopsCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *ShopsCreate) SetNillableUpdatedAt(t *time.Time) *ShopsCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *ShopsCreate) SetDeletedAt(t time.Time) *ShopsCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *ShopsCreate) SetNillableDeletedAt(t *time.Time) *ShopsCreate {
	if t != nil {
		sc.SetDeletedAt(*t)
	}
	return sc
}

// Mutation returns the ShopsMutation object of the builder.
func (sc *ShopsCreate) Mutation() *ShopsMutation {
	return sc.mutation
}

// Save creates the Shops in the database.
func (sc *ShopsCreate) Save(ctx context.Context) (*Shops, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ShopsCreate) SaveX(ctx context.Context) *Shops {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ShopsCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ShopsCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ShopsCreate) defaults() {
	if _, ok := sc.mutation.EmailVerified(); !ok {
		v := shops.DefaultEmailVerified
		sc.mutation.SetEmailVerified(v)
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := shops.DefaultCreatedAt
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := shops.DefaultUpdatedAt
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.DeletedAt(); !ok {
		v := shops.DefaultDeletedAt
		sc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ShopsCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Shops.name"`)}
	}
	if _, ok := sc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Shops.address"`)}
	}
	if _, ok := sc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Shops.email"`)}
	}
	if v, ok := sc.mutation.Email(); ok {
		if err := shops.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Shops.email": %w`, err)}
		}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Shops.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Shops.updated_at"`)}
	}
	return nil
}

func (sc *ShopsCreate) sqlSave(ctx context.Context) (*Shops, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ShopsCreate) createSpec() (*Shops, *sqlgraph.CreateSpec) {
	var (
		_node = &Shops{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(shops.Table, sqlgraph.NewFieldSpec(shops.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(shops.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := sc.mutation.Address(); ok {
		_spec.SetField(shops.FieldAddress, field.TypeString, value)
		_node.Address = &value
	}
	if value, ok := sc.mutation.Email(); ok {
		_spec.SetField(shops.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := sc.mutation.EmailVerified(); ok {
		_spec.SetField(shops.FieldEmailVerified, field.TypeTime, value)
		_node.EmailVerified = &value
	}
	if value, ok := sc.mutation.Password(); ok {
		_spec.SetField(shops.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := sc.mutation.RememberToken(); ok {
		_spec.SetField(shops.FieldRememberToken, field.TypeString, value)
		_node.RememberToken = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(shops.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(shops.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.SetField(shops.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// ShopsCreateBulk is the builder for creating many Shops entities in bulk.
type ShopsCreateBulk struct {
	config
	builders []*ShopsCreate
}

// Save creates the Shops entities in the database.
func (scb *ShopsCreateBulk) Save(ctx context.Context) ([]*Shops, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Shops, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShopsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ShopsCreateBulk) SaveX(ctx context.Context) []*Shops {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ShopsCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ShopsCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
