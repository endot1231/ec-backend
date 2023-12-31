// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/endot1231/ec-backend/ent/admins"
)

// AdminsCreate is the builder for creating a Admins entity.
type AdminsCreate struct {
	config
	mutation *AdminsMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AdminsCreate) SetName(s string) *AdminsCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetEmail sets the "email" field.
func (ac *AdminsCreate) SetEmail(s string) *AdminsCreate {
	ac.mutation.SetEmail(s)
	return ac
}

// SetEmailVerified sets the "email_verified" field.
func (ac *AdminsCreate) SetEmailVerified(t time.Time) *AdminsCreate {
	ac.mutation.SetEmailVerified(t)
	return ac
}

// SetNillableEmailVerified sets the "email_verified" field if the given value is not nil.
func (ac *AdminsCreate) SetNillableEmailVerified(t *time.Time) *AdminsCreate {
	if t != nil {
		ac.SetEmailVerified(*t)
	}
	return ac
}

// SetPassword sets the "password" field.
func (ac *AdminsCreate) SetPassword(s string) *AdminsCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetRememberToken sets the "remember_token" field.
func (ac *AdminsCreate) SetRememberToken(s string) *AdminsCreate {
	ac.mutation.SetRememberToken(s)
	return ac
}

// SetNillableRememberToken sets the "remember_token" field if the given value is not nil.
func (ac *AdminsCreate) SetNillableRememberToken(s *string) *AdminsCreate {
	if s != nil {
		ac.SetRememberToken(*s)
	}
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AdminsCreate) SetCreatedAt(t time.Time) *AdminsCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AdminsCreate) SetNillableCreatedAt(t *time.Time) *AdminsCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AdminsCreate) SetUpdatedAt(t time.Time) *AdminsCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AdminsCreate) SetNillableUpdatedAt(t *time.Time) *AdminsCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AdminsCreate) SetDeletedAt(t time.Time) *AdminsCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AdminsCreate) SetNillableDeletedAt(t *time.Time) *AdminsCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// Mutation returns the AdminsMutation object of the builder.
func (ac *AdminsCreate) Mutation() *AdminsMutation {
	return ac.mutation
}

// Save creates the Admins in the database.
func (ac *AdminsCreate) Save(ctx context.Context) (*Admins, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AdminsCreate) SaveX(ctx context.Context) *Admins {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AdminsCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AdminsCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AdminsCreate) defaults() {
	if _, ok := ac.mutation.EmailVerified(); !ok {
		v := admins.DefaultEmailVerified
		ac.mutation.SetEmailVerified(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := admins.DefaultCreatedAt
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := admins.DefaultUpdatedAt
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.DeletedAt(); !ok {
		v := admins.DefaultDeletedAt
		ac.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AdminsCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Admins.name"`)}
	}
	if _, ok := ac.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Admins.email"`)}
	}
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Admins.password"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Admins.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Admins.updated_at"`)}
	}
	return nil
}

func (ac *AdminsCreate) sqlSave(ctx context.Context) (*Admins, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AdminsCreate) createSpec() (*Admins, *sqlgraph.CreateSpec) {
	var (
		_node = &Admins{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(admins.Table, sqlgraph.NewFieldSpec(admins.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(admins.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Email(); ok {
		_spec.SetField(admins.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := ac.mutation.EmailVerified(); ok {
		_spec.SetField(admins.FieldEmailVerified, field.TypeTime, value)
		_node.EmailVerified = &value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.SetField(admins.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := ac.mutation.RememberToken(); ok {
		_spec.SetField(admins.FieldRememberToken, field.TypeString, value)
		_node.RememberToken = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(admins.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(admins.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(admins.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// AdminsCreateBulk is the builder for creating many Admins entities in bulk.
type AdminsCreateBulk struct {
	config
	builders []*AdminsCreate
}

// Save creates the Admins entities in the database.
func (acb *AdminsCreateBulk) Save(ctx context.Context) ([]*Admins, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Admins, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminsMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AdminsCreateBulk) SaveX(ctx context.Context) []*Admins {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AdminsCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AdminsCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
