// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/endot1231/ec-backend/ent/predicate"
	"github.com/endot1231/ec-backend/ent/productstock"
)

// ProductStockDelete is the builder for deleting a ProductStock entity.
type ProductStockDelete struct {
	config
	hooks    []Hook
	mutation *ProductStockMutation
}

// Where appends a list predicates to the ProductStockDelete builder.
func (psd *ProductStockDelete) Where(ps ...predicate.ProductStock) *ProductStockDelete {
	psd.mutation.Where(ps...)
	return psd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (psd *ProductStockDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, psd.sqlExec, psd.mutation, psd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (psd *ProductStockDelete) ExecX(ctx context.Context) int {
	n, err := psd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (psd *ProductStockDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(productstock.Table, sqlgraph.NewFieldSpec(productstock.FieldID, field.TypeInt))
	if ps := psd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, psd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	psd.mutation.done = true
	return affected, err
}

// ProductStockDeleteOne is the builder for deleting a single ProductStock entity.
type ProductStockDeleteOne struct {
	psd *ProductStockDelete
}

// Where appends a list predicates to the ProductStockDelete builder.
func (psdo *ProductStockDeleteOne) Where(ps ...predicate.ProductStock) *ProductStockDeleteOne {
	psdo.psd.mutation.Where(ps...)
	return psdo
}

// Exec executes the deletion query.
func (psdo *ProductStockDeleteOne) Exec(ctx context.Context) error {
	n, err := psdo.psd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productstock.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (psdo *ProductStockDeleteOne) ExecX(ctx context.Context) {
	if err := psdo.Exec(ctx); err != nil {
		panic(err)
	}
}
