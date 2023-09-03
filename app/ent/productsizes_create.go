// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/endot1231/ec-backend/ent/productsizes"
)

// ProductSizesCreate is the builder for creating a ProductSizes entity.
type ProductSizesCreate struct {
	config
	mutation *ProductSizesMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (psc *ProductSizesCreate) SetType(s string) *ProductSizesCreate {
	psc.mutation.SetType(s)
	return psc
}

// SetSize sets the "size" field.
func (psc *ProductSizesCreate) SetSize(s string) *ProductSizesCreate {
	psc.mutation.SetSize(s)
	return psc
}

// SetCreatedAt sets the "created_at" field.
func (psc *ProductSizesCreate) SetCreatedAt(t time.Time) *ProductSizesCreate {
	psc.mutation.SetCreatedAt(t)
	return psc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (psc *ProductSizesCreate) SetNillableCreatedAt(t *time.Time) *ProductSizesCreate {
	if t != nil {
		psc.SetCreatedAt(*t)
	}
	return psc
}

// SetUpdatedAt sets the "updated_at" field.
func (psc *ProductSizesCreate) SetUpdatedAt(t time.Time) *ProductSizesCreate {
	psc.mutation.SetUpdatedAt(t)
	return psc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (psc *ProductSizesCreate) SetNillableUpdatedAt(t *time.Time) *ProductSizesCreate {
	if t != nil {
		psc.SetUpdatedAt(*t)
	}
	return psc
}

// SetDeletedAt sets the "deleted_at" field.
func (psc *ProductSizesCreate) SetDeletedAt(t time.Time) *ProductSizesCreate {
	psc.mutation.SetDeletedAt(t)
	return psc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (psc *ProductSizesCreate) SetNillableDeletedAt(t *time.Time) *ProductSizesCreate {
	if t != nil {
		psc.SetDeletedAt(*t)
	}
	return psc
}

// Mutation returns the ProductSizesMutation object of the builder.
func (psc *ProductSizesCreate) Mutation() *ProductSizesMutation {
	return psc.mutation
}

// Save creates the ProductSizes in the database.
func (psc *ProductSizesCreate) Save(ctx context.Context) (*ProductSizes, error) {
	psc.defaults()
	return withHooks(ctx, psc.sqlSave, psc.mutation, psc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (psc *ProductSizesCreate) SaveX(ctx context.Context) *ProductSizes {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *ProductSizesCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *ProductSizesCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psc *ProductSizesCreate) defaults() {
	if _, ok := psc.mutation.CreatedAt(); !ok {
		v := productsizes.DefaultCreatedAt
		psc.mutation.SetCreatedAt(v)
	}
	if _, ok := psc.mutation.UpdatedAt(); !ok {
		v := productsizes.DefaultUpdatedAt
		psc.mutation.SetUpdatedAt(v)
	}
	if _, ok := psc.mutation.DeletedAt(); !ok {
		v := productsizes.DefaultDeletedAt
		psc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *ProductSizesCreate) check() error {
	if _, ok := psc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "ProductSizes.type"`)}
	}
	if _, ok := psc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "ProductSizes.size"`)}
	}
	if _, ok := psc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProductSizes.created_at"`)}
	}
	if _, ok := psc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProductSizes.updated_at"`)}
	}
	return nil
}

func (psc *ProductSizesCreate) sqlSave(ctx context.Context) (*ProductSizes, error) {
	if err := psc.check(); err != nil {
		return nil, err
	}
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	psc.mutation.id = &_node.ID
	psc.mutation.done = true
	return _node, nil
}

func (psc *ProductSizesCreate) createSpec() (*ProductSizes, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductSizes{config: psc.config}
		_spec = sqlgraph.NewCreateSpec(productsizes.Table, sqlgraph.NewFieldSpec(productsizes.FieldID, field.TypeInt))
	)
	if value, ok := psc.mutation.GetType(); ok {
		_spec.SetField(productsizes.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := psc.mutation.Size(); ok {
		_spec.SetField(productsizes.FieldSize, field.TypeString, value)
		_node.Size = value
	}
	if value, ok := psc.mutation.CreatedAt(); ok {
		_spec.SetField(productsizes.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := psc.mutation.UpdatedAt(); ok {
		_spec.SetField(productsizes.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := psc.mutation.DeletedAt(); ok {
		_spec.SetField(productsizes.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// ProductSizesCreateBulk is the builder for creating many ProductSizes entities in bulk.
type ProductSizesCreateBulk struct {
	config
	builders []*ProductSizesCreate
}

// Save creates the ProductSizes entities in the database.
func (pscb *ProductSizesCreateBulk) Save(ctx context.Context) ([]*ProductSizes, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*ProductSizes, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductSizesMutation)
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
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *ProductSizesCreateBulk) SaveX(ctx context.Context) []*ProductSizes {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *ProductSizesCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *ProductSizesCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}