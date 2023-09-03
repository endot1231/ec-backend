// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/endot1231/ec-backend/ent/reviews"
	"github.com/endot1231/ec-backend/ent/users"
)

// ReviewsCreate is the builder for creating a Reviews entity.
type ReviewsCreate struct {
	config
	mutation *ReviewsMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (rc *ReviewsCreate) SetUserID(i int) *ReviewsCreate {
	rc.mutation.SetUserID(i)
	return rc
}

// SetContents sets the "contents" field.
func (rc *ReviewsCreate) SetContents(s string) *ReviewsCreate {
	rc.mutation.SetContents(s)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *ReviewsCreate) SetCreatedAt(t time.Time) *ReviewsCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *ReviewsCreate) SetNillableCreatedAt(t *time.Time) *ReviewsCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *ReviewsCreate) SetUpdatedAt(t time.Time) *ReviewsCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *ReviewsCreate) SetNillableUpdatedAt(t *time.Time) *ReviewsCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetDeletedAt sets the "deleted_at" field.
func (rc *ReviewsCreate) SetDeletedAt(t time.Time) *ReviewsCreate {
	rc.mutation.SetDeletedAt(t)
	return rc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rc *ReviewsCreate) SetNillableDeletedAt(t *time.Time) *ReviewsCreate {
	if t != nil {
		rc.SetDeletedAt(*t)
	}
	return rc
}

// AddUserIDs adds the "user" edge to the Users entity by IDs.
func (rc *ReviewsCreate) AddUserIDs(ids ...int) *ReviewsCreate {
	rc.mutation.AddUserIDs(ids...)
	return rc
}

// AddUser adds the "user" edges to the Users entity.
func (rc *ReviewsCreate) AddUser(u ...*Users) *ReviewsCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rc.AddUserIDs(ids...)
}

// Mutation returns the ReviewsMutation object of the builder.
func (rc *ReviewsCreate) Mutation() *ReviewsMutation {
	return rc.mutation
}

// Save creates the Reviews in the database.
func (rc *ReviewsCreate) Save(ctx context.Context) (*Reviews, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReviewsCreate) SaveX(ctx context.Context) *Reviews {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReviewsCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReviewsCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReviewsCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := reviews.DefaultCreatedAt
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := reviews.DefaultUpdatedAt
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.DeletedAt(); !ok {
		v := reviews.DefaultDeletedAt
		rc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReviewsCreate) check() error {
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Reviews.user_id"`)}
	}
	if _, ok := rc.mutation.Contents(); !ok {
		return &ValidationError{Name: "contents", err: errors.New(`ent: missing required field "Reviews.contents"`)}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Reviews.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Reviews.updated_at"`)}
	}
	return nil
}

func (rc *ReviewsCreate) sqlSave(ctx context.Context) (*Reviews, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReviewsCreate) createSpec() (*Reviews, *sqlgraph.CreateSpec) {
	var (
		_node = &Reviews{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(reviews.Table, sqlgraph.NewFieldSpec(reviews.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.UserID(); ok {
		_spec.SetField(reviews.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := rc.mutation.Contents(); ok {
		_spec.SetField(reviews.FieldContents, field.TypeString, value)
		_node.Contents = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(reviews.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(reviews.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.DeletedAt(); ok {
		_spec.SetField(reviews.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   reviews.UserTable,
			Columns: reviews.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReviewsCreateBulk is the builder for creating many Reviews entities in bulk.
type ReviewsCreateBulk struct {
	config
	builders []*ReviewsCreate
}

// Save creates the Reviews entities in the database.
func (rcb *ReviewsCreateBulk) Save(ctx context.Context) ([]*Reviews, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Reviews, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReviewsMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReviewsCreateBulk) SaveX(ctx context.Context) []*Reviews {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReviewsCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReviewsCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
