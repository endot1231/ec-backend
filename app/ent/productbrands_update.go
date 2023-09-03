// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/endot1231/ec-backend/ent/predicate"
	"github.com/endot1231/ec-backend/ent/productbrands"
)

// ProductBrandsUpdate is the builder for updating ProductBrands entities.
type ProductBrandsUpdate struct {
	config
	hooks    []Hook
	mutation *ProductBrandsMutation
}

// Where appends a list predicates to the ProductBrandsUpdate builder.
func (pbu *ProductBrandsUpdate) Where(ps ...predicate.ProductBrands) *ProductBrandsUpdate {
	pbu.mutation.Where(ps...)
	return pbu
}

// SetName sets the "name" field.
func (pbu *ProductBrandsUpdate) SetName(s string) *ProductBrandsUpdate {
	pbu.mutation.SetName(s)
	return pbu
}

// SetCreatedAt sets the "created_at" field.
func (pbu *ProductBrandsUpdate) SetCreatedAt(t time.Time) *ProductBrandsUpdate {
	pbu.mutation.SetCreatedAt(t)
	return pbu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pbu *ProductBrandsUpdate) SetNillableCreatedAt(t *time.Time) *ProductBrandsUpdate {
	if t != nil {
		pbu.SetCreatedAt(*t)
	}
	return pbu
}

// SetUpdatedAt sets the "updated_at" field.
func (pbu *ProductBrandsUpdate) SetUpdatedAt(t time.Time) *ProductBrandsUpdate {
	pbu.mutation.SetUpdatedAt(t)
	return pbu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pbu *ProductBrandsUpdate) SetNillableUpdatedAt(t *time.Time) *ProductBrandsUpdate {
	if t != nil {
		pbu.SetUpdatedAt(*t)
	}
	return pbu
}

// SetDeletedAt sets the "deleted_at" field.
func (pbu *ProductBrandsUpdate) SetDeletedAt(t time.Time) *ProductBrandsUpdate {
	pbu.mutation.SetDeletedAt(t)
	return pbu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pbu *ProductBrandsUpdate) SetNillableDeletedAt(t *time.Time) *ProductBrandsUpdate {
	if t != nil {
		pbu.SetDeletedAt(*t)
	}
	return pbu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pbu *ProductBrandsUpdate) ClearDeletedAt() *ProductBrandsUpdate {
	pbu.mutation.ClearDeletedAt()
	return pbu
}

// Mutation returns the ProductBrandsMutation object of the builder.
func (pbu *ProductBrandsUpdate) Mutation() *ProductBrandsMutation {
	return pbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pbu *ProductBrandsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pbu.sqlSave, pbu.mutation, pbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pbu *ProductBrandsUpdate) SaveX(ctx context.Context) int {
	affected, err := pbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pbu *ProductBrandsUpdate) Exec(ctx context.Context) error {
	_, err := pbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbu *ProductBrandsUpdate) ExecX(ctx context.Context) {
	if err := pbu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pbu *ProductBrandsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(productbrands.Table, productbrands.Columns, sqlgraph.NewFieldSpec(productbrands.FieldID, field.TypeInt))
	if ps := pbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pbu.mutation.Name(); ok {
		_spec.SetField(productbrands.FieldName, field.TypeString, value)
	}
	if value, ok := pbu.mutation.CreatedAt(); ok {
		_spec.SetField(productbrands.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pbu.mutation.UpdatedAt(); ok {
		_spec.SetField(productbrands.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pbu.mutation.DeletedAt(); ok {
		_spec.SetField(productbrands.FieldDeletedAt, field.TypeTime, value)
	}
	if pbu.mutation.DeletedAtCleared() {
		_spec.ClearField(productbrands.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productbrands.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pbu.mutation.done = true
	return n, nil
}

// ProductBrandsUpdateOne is the builder for updating a single ProductBrands entity.
type ProductBrandsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductBrandsMutation
}

// SetName sets the "name" field.
func (pbuo *ProductBrandsUpdateOne) SetName(s string) *ProductBrandsUpdateOne {
	pbuo.mutation.SetName(s)
	return pbuo
}

// SetCreatedAt sets the "created_at" field.
func (pbuo *ProductBrandsUpdateOne) SetCreatedAt(t time.Time) *ProductBrandsUpdateOne {
	pbuo.mutation.SetCreatedAt(t)
	return pbuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pbuo *ProductBrandsUpdateOne) SetNillableCreatedAt(t *time.Time) *ProductBrandsUpdateOne {
	if t != nil {
		pbuo.SetCreatedAt(*t)
	}
	return pbuo
}

// SetUpdatedAt sets the "updated_at" field.
func (pbuo *ProductBrandsUpdateOne) SetUpdatedAt(t time.Time) *ProductBrandsUpdateOne {
	pbuo.mutation.SetUpdatedAt(t)
	return pbuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pbuo *ProductBrandsUpdateOne) SetNillableUpdatedAt(t *time.Time) *ProductBrandsUpdateOne {
	if t != nil {
		pbuo.SetUpdatedAt(*t)
	}
	return pbuo
}

// SetDeletedAt sets the "deleted_at" field.
func (pbuo *ProductBrandsUpdateOne) SetDeletedAt(t time.Time) *ProductBrandsUpdateOne {
	pbuo.mutation.SetDeletedAt(t)
	return pbuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pbuo *ProductBrandsUpdateOne) SetNillableDeletedAt(t *time.Time) *ProductBrandsUpdateOne {
	if t != nil {
		pbuo.SetDeletedAt(*t)
	}
	return pbuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pbuo *ProductBrandsUpdateOne) ClearDeletedAt() *ProductBrandsUpdateOne {
	pbuo.mutation.ClearDeletedAt()
	return pbuo
}

// Mutation returns the ProductBrandsMutation object of the builder.
func (pbuo *ProductBrandsUpdateOne) Mutation() *ProductBrandsMutation {
	return pbuo.mutation
}

// Where appends a list predicates to the ProductBrandsUpdate builder.
func (pbuo *ProductBrandsUpdateOne) Where(ps ...predicate.ProductBrands) *ProductBrandsUpdateOne {
	pbuo.mutation.Where(ps...)
	return pbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pbuo *ProductBrandsUpdateOne) Select(field string, fields ...string) *ProductBrandsUpdateOne {
	pbuo.fields = append([]string{field}, fields...)
	return pbuo
}

// Save executes the query and returns the updated ProductBrands entity.
func (pbuo *ProductBrandsUpdateOne) Save(ctx context.Context) (*ProductBrands, error) {
	return withHooks(ctx, pbuo.sqlSave, pbuo.mutation, pbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pbuo *ProductBrandsUpdateOne) SaveX(ctx context.Context) *ProductBrands {
	node, err := pbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pbuo *ProductBrandsUpdateOne) Exec(ctx context.Context) error {
	_, err := pbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbuo *ProductBrandsUpdateOne) ExecX(ctx context.Context) {
	if err := pbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pbuo *ProductBrandsUpdateOne) sqlSave(ctx context.Context) (_node *ProductBrands, err error) {
	_spec := sqlgraph.NewUpdateSpec(productbrands.Table, productbrands.Columns, sqlgraph.NewFieldSpec(productbrands.FieldID, field.TypeInt))
	id, ok := pbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProductBrands.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productbrands.FieldID)
		for _, f := range fields {
			if !productbrands.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != productbrands.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pbuo.mutation.Name(); ok {
		_spec.SetField(productbrands.FieldName, field.TypeString, value)
	}
	if value, ok := pbuo.mutation.CreatedAt(); ok {
		_spec.SetField(productbrands.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pbuo.mutation.UpdatedAt(); ok {
		_spec.SetField(productbrands.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pbuo.mutation.DeletedAt(); ok {
		_spec.SetField(productbrands.FieldDeletedAt, field.TypeTime, value)
	}
	if pbuo.mutation.DeletedAtCleared() {
		_spec.ClearField(productbrands.FieldDeletedAt, field.TypeTime)
	}
	_node = &ProductBrands{config: pbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productbrands.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pbuo.mutation.done = true
	return _node, nil
}