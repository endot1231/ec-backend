// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/endot1231/ec-backend/ent/predicate"
	"github.com/endot1231/ec-backend/ent/productcolors"
)

// ProductColorsQuery is the builder for querying ProductColors entities.
type ProductColorsQuery struct {
	config
	ctx        *QueryContext
	order      []productcolors.OrderOption
	inters     []Interceptor
	predicates []predicate.ProductColors
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductColorsQuery builder.
func (pcq *ProductColorsQuery) Where(ps ...predicate.ProductColors) *ProductColorsQuery {
	pcq.predicates = append(pcq.predicates, ps...)
	return pcq
}

// Limit the number of records to be returned by this query.
func (pcq *ProductColorsQuery) Limit(limit int) *ProductColorsQuery {
	pcq.ctx.Limit = &limit
	return pcq
}

// Offset to start from.
func (pcq *ProductColorsQuery) Offset(offset int) *ProductColorsQuery {
	pcq.ctx.Offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *ProductColorsQuery) Unique(unique bool) *ProductColorsQuery {
	pcq.ctx.Unique = &unique
	return pcq
}

// Order specifies how the records should be ordered.
func (pcq *ProductColorsQuery) Order(o ...productcolors.OrderOption) *ProductColorsQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// First returns the first ProductColors entity from the query.
// Returns a *NotFoundError when no ProductColors was found.
func (pcq *ProductColorsQuery) First(ctx context.Context) (*ProductColors, error) {
	nodes, err := pcq.Limit(1).All(setContextOp(ctx, pcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productcolors.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcq *ProductColorsQuery) FirstX(ctx context.Context) *ProductColors {
	node, err := pcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductColors ID from the query.
// Returns a *NotFoundError when no ProductColors ID was found.
func (pcq *ProductColorsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pcq.Limit(1).IDs(setContextOp(ctx, pcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productcolors.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcq *ProductColorsQuery) FirstIDX(ctx context.Context) int {
	id, err := pcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductColors entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProductColors entity is found.
// Returns a *NotFoundError when no ProductColors entities are found.
func (pcq *ProductColorsQuery) Only(ctx context.Context) (*ProductColors, error) {
	nodes, err := pcq.Limit(2).All(setContextOp(ctx, pcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productcolors.Label}
	default:
		return nil, &NotSingularError{productcolors.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcq *ProductColorsQuery) OnlyX(ctx context.Context) *ProductColors {
	node, err := pcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductColors ID in the query.
// Returns a *NotSingularError when more than one ProductColors ID is found.
// Returns a *NotFoundError when no entities are found.
func (pcq *ProductColorsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pcq.Limit(2).IDs(setContextOp(ctx, pcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productcolors.Label}
	default:
		err = &NotSingularError{productcolors.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcq *ProductColorsQuery) OnlyIDX(ctx context.Context) int {
	id, err := pcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductColorsSlice.
func (pcq *ProductColorsQuery) All(ctx context.Context) ([]*ProductColors, error) {
	ctx = setContextOp(ctx, pcq.ctx, "All")
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProductColors, *ProductColorsQuery]()
	return withInterceptors[[]*ProductColors](ctx, pcq, qr, pcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pcq *ProductColorsQuery) AllX(ctx context.Context) []*ProductColors {
	nodes, err := pcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductColors IDs.
func (pcq *ProductColorsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pcq.ctx.Unique == nil && pcq.path != nil {
		pcq.Unique(true)
	}
	ctx = setContextOp(ctx, pcq.ctx, "IDs")
	if err = pcq.Select(productcolors.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcq *ProductColorsQuery) IDsX(ctx context.Context) []int {
	ids, err := pcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcq *ProductColorsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Count")
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pcq, querierCount[*ProductColorsQuery](), pcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pcq *ProductColorsQuery) CountX(ctx context.Context) int {
	count, err := pcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcq *ProductColorsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Exist")
	switch _, err := pcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pcq *ProductColorsQuery) ExistX(ctx context.Context) bool {
	exist, err := pcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductColorsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcq *ProductColorsQuery) Clone() *ProductColorsQuery {
	if pcq == nil {
		return nil
	}
	return &ProductColorsQuery{
		config:     pcq.config,
		ctx:        pcq.ctx.Clone(),
		order:      append([]productcolors.OrderOption{}, pcq.order...),
		inters:     append([]Interceptor{}, pcq.inters...),
		predicates: append([]predicate.ProductColors{}, pcq.predicates...),
		// clone intermediate query.
		sql:  pcq.sql.Clone(),
		path: pcq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProductColors.Query().
//		GroupBy(productcolors.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pcq *ProductColorsQuery) GroupBy(field string, fields ...string) *ProductColorsGroupBy {
	pcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProductColorsGroupBy{build: pcq}
	grbuild.flds = &pcq.ctx.Fields
	grbuild.label = productcolors.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ProductColors.Query().
//		Select(productcolors.FieldName).
//		Scan(ctx, &v)
func (pcq *ProductColorsQuery) Select(fields ...string) *ProductColorsSelect {
	pcq.ctx.Fields = append(pcq.ctx.Fields, fields...)
	sbuild := &ProductColorsSelect{ProductColorsQuery: pcq}
	sbuild.label = productcolors.Label
	sbuild.flds, sbuild.scan = &pcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProductColorsSelect configured with the given aggregations.
func (pcq *ProductColorsQuery) Aggregate(fns ...AggregateFunc) *ProductColorsSelect {
	return pcq.Select().Aggregate(fns...)
}

func (pcq *ProductColorsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pcq); err != nil {
				return err
			}
		}
	}
	for _, f := range pcq.ctx.Fields {
		if !productcolors.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcq.path != nil {
		prev, err := pcq.path(ctx)
		if err != nil {
			return err
		}
		pcq.sql = prev
	}
	return nil
}

func (pcq *ProductColorsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProductColors, error) {
	var (
		nodes = []*ProductColors{}
		_spec = pcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProductColors).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProductColors{config: pcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pcq *ProductColorsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcq.querySpec()
	_spec.Node.Columns = pcq.ctx.Fields
	if len(pcq.ctx.Fields) > 0 {
		_spec.Unique = pcq.ctx.Unique != nil && *pcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *ProductColorsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(productcolors.Table, productcolors.Columns, sqlgraph.NewFieldSpec(productcolors.FieldID, field.TypeInt))
	_spec.From = pcq.sql
	if unique := pcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pcq.path != nil {
		_spec.Unique = true
	}
	if fields := pcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productcolors.FieldID)
		for i := range fields {
			if fields[i] != productcolors.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcq *ProductColorsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcq.driver.Dialect())
	t1 := builder.Table(productcolors.Table)
	columns := pcq.ctx.Fields
	if len(columns) == 0 {
		columns = productcolors.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcq.ctx.Unique != nil && *pcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductColorsGroupBy is the group-by builder for ProductColors entities.
type ProductColorsGroupBy struct {
	selector
	build *ProductColorsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *ProductColorsGroupBy) Aggregate(fns ...AggregateFunc) *ProductColorsGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the selector query and scans the result into the given value.
func (pcgb *ProductColorsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcgb.build.ctx, "GroupBy")
	if err := pcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductColorsQuery, *ProductColorsGroupBy](ctx, pcgb.build, pcgb, pcgb.build.inters, v)
}

func (pcgb *ProductColorsGroupBy) sqlScan(ctx context.Context, root *ProductColorsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pcgb.flds)+len(pcgb.fns))
		for _, f := range *pcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProductColorsSelect is the builder for selecting fields of ProductColors entities.
type ProductColorsSelect struct {
	*ProductColorsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pcs *ProductColorsSelect) Aggregate(fns ...AggregateFunc) *ProductColorsSelect {
	pcs.fns = append(pcs.fns, fns...)
	return pcs
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *ProductColorsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcs.ctx, "Select")
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductColorsQuery, *ProductColorsSelect](ctx, pcs.ProductColorsQuery, pcs, pcs.inters, v)
}

func (pcs *ProductColorsSelect) sqlScan(ctx context.Context, root *ProductColorsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pcs.fns))
	for _, fn := range pcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}