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
	"github.com/endot1231/ec-backend/ent/productcategories"
)

// ProductCategoriesUpdate is the builder for updating ProductCategories entities.
type ProductCategoriesUpdate struct {
	config
	hooks    []Hook
	mutation *ProductCategoriesMutation
}

// Where appends a list predicates to the ProductCategoriesUpdate builder.
func (pcu *ProductCategoriesUpdate) Where(ps ...predicate.ProductCategories) *ProductCategoriesUpdate {
	pcu.mutation.Where(ps...)
	return pcu
}

// SetName sets the "name" field.
func (pcu *ProductCategoriesUpdate) SetName(s string) *ProductCategoriesUpdate {
	pcu.mutation.SetName(s)
	return pcu
}

// SetCreatedAt sets the "created_at" field.
func (pcu *ProductCategoriesUpdate) SetCreatedAt(t time.Time) *ProductCategoriesUpdate {
	pcu.mutation.SetCreatedAt(t)
	return pcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pcu *ProductCategoriesUpdate) SetNillableCreatedAt(t *time.Time) *ProductCategoriesUpdate {
	if t != nil {
		pcu.SetCreatedAt(*t)
	}
	return pcu
}

// SetUpdatedAt sets the "updated_at" field.
func (pcu *ProductCategoriesUpdate) SetUpdatedAt(t time.Time) *ProductCategoriesUpdate {
	pcu.mutation.SetUpdatedAt(t)
	return pcu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pcu *ProductCategoriesUpdate) SetNillableUpdatedAt(t *time.Time) *ProductCategoriesUpdate {
	if t != nil {
		pcu.SetUpdatedAt(*t)
	}
	return pcu
}

// SetDeletedAt sets the "deleted_at" field.
func (pcu *ProductCategoriesUpdate) SetDeletedAt(t time.Time) *ProductCategoriesUpdate {
	pcu.mutation.SetDeletedAt(t)
	return pcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pcu *ProductCategoriesUpdate) SetNillableDeletedAt(t *time.Time) *ProductCategoriesUpdate {
	if t != nil {
		pcu.SetDeletedAt(*t)
	}
	return pcu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pcu *ProductCategoriesUpdate) ClearDeletedAt() *ProductCategoriesUpdate {
	pcu.mutation.ClearDeletedAt()
	return pcu
}

// Mutation returns the ProductCategoriesMutation object of the builder.
func (pcu *ProductCategoriesUpdate) Mutation() *ProductCategoriesMutation {
	return pcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *ProductCategoriesUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pcu.sqlSave, pcu.mutation, pcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcu *ProductCategoriesUpdate) SaveX(ctx context.Context) int {
	affected, err := pcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcu *ProductCategoriesUpdate) Exec(ctx context.Context) error {
	_, err := pcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcu *ProductCategoriesUpdate) ExecX(ctx context.Context) {
	if err := pcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pcu *ProductCategoriesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(productcategories.Table, productcategories.Columns, sqlgraph.NewFieldSpec(productcategories.FieldID, field.TypeInt))
	if ps := pcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcu.mutation.Name(); ok {
		_spec.SetField(productcategories.FieldName, field.TypeString, value)
	}
	if value, ok := pcu.mutation.CreatedAt(); ok {
		_spec.SetField(productcategories.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pcu.mutation.UpdatedAt(); ok {
		_spec.SetField(productcategories.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pcu.mutation.DeletedAt(); ok {
		_spec.SetField(productcategories.FieldDeletedAt, field.TypeTime, value)
	}
	if pcu.mutation.DeletedAtCleared() {
		_spec.ClearField(productcategories.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productcategories.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pcu.mutation.done = true
	return n, nil
}

// ProductCategoriesUpdateOne is the builder for updating a single ProductCategories entity.
type ProductCategoriesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductCategoriesMutation
}

// SetName sets the "name" field.
func (pcuo *ProductCategoriesUpdateOne) SetName(s string) *ProductCategoriesUpdateOne {
	pcuo.mutation.SetName(s)
	return pcuo
}

// SetCreatedAt sets the "created_at" field.
func (pcuo *ProductCategoriesUpdateOne) SetCreatedAt(t time.Time) *ProductCategoriesUpdateOne {
	pcuo.mutation.SetCreatedAt(t)
	return pcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pcuo *ProductCategoriesUpdateOne) SetNillableCreatedAt(t *time.Time) *ProductCategoriesUpdateOne {
	if t != nil {
		pcuo.SetCreatedAt(*t)
	}
	return pcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (pcuo *ProductCategoriesUpdateOne) SetUpdatedAt(t time.Time) *ProductCategoriesUpdateOne {
	pcuo.mutation.SetUpdatedAt(t)
	return pcuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pcuo *ProductCategoriesUpdateOne) SetNillableUpdatedAt(t *time.Time) *ProductCategoriesUpdateOne {
	if t != nil {
		pcuo.SetUpdatedAt(*t)
	}
	return pcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (pcuo *ProductCategoriesUpdateOne) SetDeletedAt(t time.Time) *ProductCategoriesUpdateOne {
	pcuo.mutation.SetDeletedAt(t)
	return pcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pcuo *ProductCategoriesUpdateOne) SetNillableDeletedAt(t *time.Time) *ProductCategoriesUpdateOne {
	if t != nil {
		pcuo.SetDeletedAt(*t)
	}
	return pcuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pcuo *ProductCategoriesUpdateOne) ClearDeletedAt() *ProductCategoriesUpdateOne {
	pcuo.mutation.ClearDeletedAt()
	return pcuo
}

// Mutation returns the ProductCategoriesMutation object of the builder.
func (pcuo *ProductCategoriesUpdateOne) Mutation() *ProductCategoriesMutation {
	return pcuo.mutation
}

// Where appends a list predicates to the ProductCategoriesUpdate builder.
func (pcuo *ProductCategoriesUpdateOne) Where(ps ...predicate.ProductCategories) *ProductCategoriesUpdateOne {
	pcuo.mutation.Where(ps...)
	return pcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcuo *ProductCategoriesUpdateOne) Select(field string, fields ...string) *ProductCategoriesUpdateOne {
	pcuo.fields = append([]string{field}, fields...)
	return pcuo
}

// Save executes the query and returns the updated ProductCategories entity.
func (pcuo *ProductCategoriesUpdateOne) Save(ctx context.Context) (*ProductCategories, error) {
	return withHooks(ctx, pcuo.sqlSave, pcuo.mutation, pcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcuo *ProductCategoriesUpdateOne) SaveX(ctx context.Context) *ProductCategories {
	node, err := pcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcuo *ProductCategoriesUpdateOne) Exec(ctx context.Context) error {
	_, err := pcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcuo *ProductCategoriesUpdateOne) ExecX(ctx context.Context) {
	if err := pcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pcuo *ProductCategoriesUpdateOne) sqlSave(ctx context.Context) (_node *ProductCategories, err error) {
	_spec := sqlgraph.NewUpdateSpec(productcategories.Table, productcategories.Columns, sqlgraph.NewFieldSpec(productcategories.FieldID, field.TypeInt))
	id, ok := pcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProductCategories.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productcategories.FieldID)
		for _, f := range fields {
			if !productcategories.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != productcategories.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcuo.mutation.Name(); ok {
		_spec.SetField(productcategories.FieldName, field.TypeString, value)
	}
	if value, ok := pcuo.mutation.CreatedAt(); ok {
		_spec.SetField(productcategories.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(productcategories.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pcuo.mutation.DeletedAt(); ok {
		_spec.SetField(productcategories.FieldDeletedAt, field.TypeTime, value)
	}
	if pcuo.mutation.DeletedAtCleared() {
		_spec.ClearField(productcategories.FieldDeletedAt, field.TypeTime)
	}
	_node = &ProductCategories{config: pcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productcategories.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pcuo.mutation.done = true
	return _node, nil
}
