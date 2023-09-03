// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/endot1231/ec-backend/ent/productbrands"
)

// ProductBrandsCreate is the builder for creating a ProductBrands entity.
type ProductBrandsCreate struct {
	config
	mutation *ProductBrandsMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pbc *ProductBrandsCreate) SetName(s string) *ProductBrandsCreate {
	pbc.mutation.SetName(s)
	return pbc
}

// SetCreatedAt sets the "created_at" field.
func (pbc *ProductBrandsCreate) SetCreatedAt(t time.Time) *ProductBrandsCreate {
	pbc.mutation.SetCreatedAt(t)
	return pbc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pbc *ProductBrandsCreate) SetNillableCreatedAt(t *time.Time) *ProductBrandsCreate {
	if t != nil {
		pbc.SetCreatedAt(*t)
	}
	return pbc
}

// SetUpdatedAt sets the "updated_at" field.
func (pbc *ProductBrandsCreate) SetUpdatedAt(t time.Time) *ProductBrandsCreate {
	pbc.mutation.SetUpdatedAt(t)
	return pbc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pbc *ProductBrandsCreate) SetNillableUpdatedAt(t *time.Time) *ProductBrandsCreate {
	if t != nil {
		pbc.SetUpdatedAt(*t)
	}
	return pbc
}

// SetDeletedAt sets the "deleted_at" field.
func (pbc *ProductBrandsCreate) SetDeletedAt(t time.Time) *ProductBrandsCreate {
	pbc.mutation.SetDeletedAt(t)
	return pbc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pbc *ProductBrandsCreate) SetNillableDeletedAt(t *time.Time) *ProductBrandsCreate {
	if t != nil {
		pbc.SetDeletedAt(*t)
	}
	return pbc
}

// Mutation returns the ProductBrandsMutation object of the builder.
func (pbc *ProductBrandsCreate) Mutation() *ProductBrandsMutation {
	return pbc.mutation
}

// Save creates the ProductBrands in the database.
func (pbc *ProductBrandsCreate) Save(ctx context.Context) (*ProductBrands, error) {
	pbc.defaults()
	return withHooks(ctx, pbc.sqlSave, pbc.mutation, pbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pbc *ProductBrandsCreate) SaveX(ctx context.Context) *ProductBrands {
	v, err := pbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pbc *ProductBrandsCreate) Exec(ctx context.Context) error {
	_, err := pbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbc *ProductBrandsCreate) ExecX(ctx context.Context) {
	if err := pbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pbc *ProductBrandsCreate) defaults() {
	if _, ok := pbc.mutation.CreatedAt(); !ok {
		v := productbrands.DefaultCreatedAt
		pbc.mutation.SetCreatedAt(v)
	}
	if _, ok := pbc.mutation.UpdatedAt(); !ok {
		v := productbrands.DefaultUpdatedAt
		pbc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pbc.mutation.DeletedAt(); !ok {
		v := productbrands.DefaultDeletedAt
		pbc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pbc *ProductBrandsCreate) check() error {
	if _, ok := pbc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ProductBrands.name"`)}
	}
	if _, ok := pbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProductBrands.created_at"`)}
	}
	if _, ok := pbc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProductBrands.updated_at"`)}
	}
	return nil
}

func (pbc *ProductBrandsCreate) sqlSave(ctx context.Context) (*ProductBrands, error) {
	if err := pbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pbc.mutation.id = &_node.ID
	pbc.mutation.done = true
	return _node, nil
}

func (pbc *ProductBrandsCreate) createSpec() (*ProductBrands, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductBrands{config: pbc.config}
		_spec = sqlgraph.NewCreateSpec(productbrands.Table, sqlgraph.NewFieldSpec(productbrands.FieldID, field.TypeInt))
	)
	if value, ok := pbc.mutation.Name(); ok {
		_spec.SetField(productbrands.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pbc.mutation.CreatedAt(); ok {
		_spec.SetField(productbrands.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pbc.mutation.UpdatedAt(); ok {
		_spec.SetField(productbrands.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pbc.mutation.DeletedAt(); ok {
		_spec.SetField(productbrands.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// ProductBrandsCreateBulk is the builder for creating many ProductBrands entities in bulk.
type ProductBrandsCreateBulk struct {
	config
	builders []*ProductBrandsCreate
}

// Save creates the ProductBrands entities in the database.
func (pbcb *ProductBrandsCreateBulk) Save(ctx context.Context) ([]*ProductBrands, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pbcb.builders))
	nodes := make([]*ProductBrands, len(pbcb.builders))
	mutators := make([]Mutator, len(pbcb.builders))
	for i := range pbcb.builders {
		func(i int, root context.Context) {
			builder := pbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductBrandsMutation)
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
					_, err = mutators[i+1].Mutate(root, pbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pbcb *ProductBrandsCreateBulk) SaveX(ctx context.Context) []*ProductBrands {
	v, err := pbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pbcb *ProductBrandsCreateBulk) Exec(ctx context.Context) error {
	_, err := pbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbcb *ProductBrandsCreateBulk) ExecX(ctx context.Context) {
	if err := pbcb.Exec(ctx); err != nil {
		panic(err)
	}
}