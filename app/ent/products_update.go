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
	"github.com/endot1231/ec-backend/ent/products"
)

// ProductsUpdate is the builder for updating Products entities.
type ProductsUpdate struct {
	config
	hooks    []Hook
	mutation *ProductsMutation
}

// Where appends a list predicates to the ProductsUpdate builder.
func (pu *ProductsUpdate) Where(ps ...predicate.Products) *ProductsUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProductsUpdate) SetName(s string) *ProductsUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProductsUpdate) SetDescription(s string) *ProductsUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *ProductsUpdate) SetCreatedAt(t time.Time) *ProductsUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *ProductsUpdate) SetNillableCreatedAt(t *time.Time) *ProductsUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProductsUpdate) SetUpdatedAt(t time.Time) *ProductsUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *ProductsUpdate) SetNillableUpdatedAt(t *time.Time) *ProductsUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// SetDeletedAt sets the "deleted_at" field.
func (pu *ProductsUpdate) SetDeletedAt(t time.Time) *ProductsUpdate {
	pu.mutation.SetDeletedAt(t)
	return pu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pu *ProductsUpdate) SetNillableDeletedAt(t *time.Time) *ProductsUpdate {
	if t != nil {
		pu.SetDeletedAt(*t)
	}
	return pu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pu *ProductsUpdate) ClearDeletedAt() *ProductsUpdate {
	pu.mutation.ClearDeletedAt()
	return pu
}

// Mutation returns the ProductsMutation object of the builder.
func (pu *ProductsUpdate) Mutation() *ProductsMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProductsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductsUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductsUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductsUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProductsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(products.Table, products.Columns, sqlgraph.NewFieldSpec(products.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(products.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(products.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(products.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(products.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.SetField(products.FieldDeletedAt, field.TypeTime, value)
	}
	if pu.mutation.DeletedAtCleared() {
		_spec.ClearField(products.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{products.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProductsUpdateOne is the builder for updating a single Products entity.
type ProductsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductsMutation
}

// SetName sets the "name" field.
func (puo *ProductsUpdateOne) SetName(s string) *ProductsUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProductsUpdateOne) SetDescription(s string) *ProductsUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *ProductsUpdateOne) SetCreatedAt(t time.Time) *ProductsUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *ProductsUpdateOne) SetNillableCreatedAt(t *time.Time) *ProductsUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProductsUpdateOne) SetUpdatedAt(t time.Time) *ProductsUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *ProductsUpdateOne) SetNillableUpdatedAt(t *time.Time) *ProductsUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// SetDeletedAt sets the "deleted_at" field.
func (puo *ProductsUpdateOne) SetDeletedAt(t time.Time) *ProductsUpdateOne {
	puo.mutation.SetDeletedAt(t)
	return puo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (puo *ProductsUpdateOne) SetNillableDeletedAt(t *time.Time) *ProductsUpdateOne {
	if t != nil {
		puo.SetDeletedAt(*t)
	}
	return puo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (puo *ProductsUpdateOne) ClearDeletedAt() *ProductsUpdateOne {
	puo.mutation.ClearDeletedAt()
	return puo
}

// Mutation returns the ProductsMutation object of the builder.
func (puo *ProductsUpdateOne) Mutation() *ProductsMutation {
	return puo.mutation
}

// Where appends a list predicates to the ProductsUpdate builder.
func (puo *ProductsUpdateOne) Where(ps ...predicate.Products) *ProductsUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProductsUpdateOne) Select(field string, fields ...string) *ProductsUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Products entity.
func (puo *ProductsUpdateOne) Save(ctx context.Context) (*Products, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductsUpdateOne) SaveX(ctx context.Context) *Products {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProductsUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductsUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProductsUpdateOne) sqlSave(ctx context.Context) (_node *Products, err error) {
	_spec := sqlgraph.NewUpdateSpec(products.Table, products.Columns, sqlgraph.NewFieldSpec(products.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Products.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, products.FieldID)
		for _, f := range fields {
			if !products.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != products.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(products.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(products.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(products.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(products.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.SetField(products.FieldDeletedAt, field.TypeTime, value)
	}
	if puo.mutation.DeletedAtCleared() {
		_spec.ClearField(products.FieldDeletedAt, field.TypeTime)
	}
	_node = &Products{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{products.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
