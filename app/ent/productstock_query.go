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
	"github.com/endot1231/ec-backend/ent/productstock"
)

// ProductStockQuery is the builder for querying ProductStock entities.
type ProductStockQuery struct {
	config
	ctx        *QueryContext
	order      []productstock.OrderOption
	inters     []Interceptor
	predicates []predicate.ProductStock
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductStockQuery builder.
func (psq *ProductStockQuery) Where(ps ...predicate.ProductStock) *ProductStockQuery {
	psq.predicates = append(psq.predicates, ps...)
	return psq
}

// Limit the number of records to be returned by this query.
func (psq *ProductStockQuery) Limit(limit int) *ProductStockQuery {
	psq.ctx.Limit = &limit
	return psq
}

// Offset to start from.
func (psq *ProductStockQuery) Offset(offset int) *ProductStockQuery {
	psq.ctx.Offset = &offset
	return psq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (psq *ProductStockQuery) Unique(unique bool) *ProductStockQuery {
	psq.ctx.Unique = &unique
	return psq
}

// Order specifies how the records should be ordered.
func (psq *ProductStockQuery) Order(o ...productstock.OrderOption) *ProductStockQuery {
	psq.order = append(psq.order, o...)
	return psq
}

// First returns the first ProductStock entity from the query.
// Returns a *NotFoundError when no ProductStock was found.
func (psq *ProductStockQuery) First(ctx context.Context) (*ProductStock, error) {
	nodes, err := psq.Limit(1).All(setContextOp(ctx, psq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productstock.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (psq *ProductStockQuery) FirstX(ctx context.Context) *ProductStock {
	node, err := psq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductStock ID from the query.
// Returns a *NotFoundError when no ProductStock ID was found.
func (psq *ProductStockQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(1).IDs(setContextOp(ctx, psq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productstock.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (psq *ProductStockQuery) FirstIDX(ctx context.Context) int {
	id, err := psq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductStock entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProductStock entity is found.
// Returns a *NotFoundError when no ProductStock entities are found.
func (psq *ProductStockQuery) Only(ctx context.Context) (*ProductStock, error) {
	nodes, err := psq.Limit(2).All(setContextOp(ctx, psq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productstock.Label}
	default:
		return nil, &NotSingularError{productstock.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (psq *ProductStockQuery) OnlyX(ctx context.Context) *ProductStock {
	node, err := psq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductStock ID in the query.
// Returns a *NotSingularError when more than one ProductStock ID is found.
// Returns a *NotFoundError when no entities are found.
func (psq *ProductStockQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(2).IDs(setContextOp(ctx, psq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productstock.Label}
	default:
		err = &NotSingularError{productstock.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (psq *ProductStockQuery) OnlyIDX(ctx context.Context) int {
	id, err := psq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductStocks.
func (psq *ProductStockQuery) All(ctx context.Context) ([]*ProductStock, error) {
	ctx = setContextOp(ctx, psq.ctx, "All")
	if err := psq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProductStock, *ProductStockQuery]()
	return withInterceptors[[]*ProductStock](ctx, psq, qr, psq.inters)
}

// AllX is like All, but panics if an error occurs.
func (psq *ProductStockQuery) AllX(ctx context.Context) []*ProductStock {
	nodes, err := psq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductStock IDs.
func (psq *ProductStockQuery) IDs(ctx context.Context) (ids []int, err error) {
	if psq.ctx.Unique == nil && psq.path != nil {
		psq.Unique(true)
	}
	ctx = setContextOp(ctx, psq.ctx, "IDs")
	if err = psq.Select(productstock.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (psq *ProductStockQuery) IDsX(ctx context.Context) []int {
	ids, err := psq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (psq *ProductStockQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, psq.ctx, "Count")
	if err := psq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, psq, querierCount[*ProductStockQuery](), psq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (psq *ProductStockQuery) CountX(ctx context.Context) int {
	count, err := psq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (psq *ProductStockQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, psq.ctx, "Exist")
	switch _, err := psq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (psq *ProductStockQuery) ExistX(ctx context.Context) bool {
	exist, err := psq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductStockQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (psq *ProductStockQuery) Clone() *ProductStockQuery {
	if psq == nil {
		return nil
	}
	return &ProductStockQuery{
		config:     psq.config,
		ctx:        psq.ctx.Clone(),
		order:      append([]productstock.OrderOption{}, psq.order...),
		inters:     append([]Interceptor{}, psq.inters...),
		predicates: append([]predicate.ProductStock{}, psq.predicates...),
		// clone intermediate query.
		sql:  psq.sql.Clone(),
		path: psq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProductStock.Query().
//		GroupBy(productstock.FieldDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (psq *ProductStockQuery) GroupBy(field string, fields ...string) *ProductStockGroupBy {
	psq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProductStockGroupBy{build: psq}
	grbuild.flds = &psq.ctx.Fields
	grbuild.label = productstock.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//	}
//
//	client.ProductStock.Query().
//		Select(productstock.FieldDescription).
//		Scan(ctx, &v)
func (psq *ProductStockQuery) Select(fields ...string) *ProductStockSelect {
	psq.ctx.Fields = append(psq.ctx.Fields, fields...)
	sbuild := &ProductStockSelect{ProductStockQuery: psq}
	sbuild.label = productstock.Label
	sbuild.flds, sbuild.scan = &psq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProductStockSelect configured with the given aggregations.
func (psq *ProductStockQuery) Aggregate(fns ...AggregateFunc) *ProductStockSelect {
	return psq.Select().Aggregate(fns...)
}

func (psq *ProductStockQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range psq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, psq); err != nil {
				return err
			}
		}
	}
	for _, f := range psq.ctx.Fields {
		if !productstock.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if psq.path != nil {
		prev, err := psq.path(ctx)
		if err != nil {
			return err
		}
		psq.sql = prev
	}
	return nil
}

func (psq *ProductStockQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProductStock, error) {
	var (
		nodes = []*ProductStock{}
		_spec = psq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProductStock).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProductStock{config: psq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, psq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (psq *ProductStockQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := psq.querySpec()
	_spec.Node.Columns = psq.ctx.Fields
	if len(psq.ctx.Fields) > 0 {
		_spec.Unique = psq.ctx.Unique != nil && *psq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, psq.driver, _spec)
}

func (psq *ProductStockQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(productstock.Table, productstock.Columns, sqlgraph.NewFieldSpec(productstock.FieldID, field.TypeInt))
	_spec.From = psq.sql
	if unique := psq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if psq.path != nil {
		_spec.Unique = true
	}
	if fields := psq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productstock.FieldID)
		for i := range fields {
			if fields[i] != productstock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := psq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := psq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := psq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := psq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (psq *ProductStockQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(psq.driver.Dialect())
	t1 := builder.Table(productstock.Table)
	columns := psq.ctx.Fields
	if len(columns) == 0 {
		columns = productstock.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if psq.sql != nil {
		selector = psq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if psq.ctx.Unique != nil && *psq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range psq.predicates {
		p(selector)
	}
	for _, p := range psq.order {
		p(selector)
	}
	if offset := psq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := psq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductStockGroupBy is the group-by builder for ProductStock entities.
type ProductStockGroupBy struct {
	selector
	build *ProductStockQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (psgb *ProductStockGroupBy) Aggregate(fns ...AggregateFunc) *ProductStockGroupBy {
	psgb.fns = append(psgb.fns, fns...)
	return psgb
}

// Scan applies the selector query and scans the result into the given value.
func (psgb *ProductStockGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psgb.build.ctx, "GroupBy")
	if err := psgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductStockQuery, *ProductStockGroupBy](ctx, psgb.build, psgb, psgb.build.inters, v)
}

func (psgb *ProductStockGroupBy) sqlScan(ctx context.Context, root *ProductStockQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(psgb.fns))
	for _, fn := range psgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*psgb.flds)+len(psgb.fns))
		for _, f := range *psgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*psgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProductStockSelect is the builder for selecting fields of ProductStock entities.
type ProductStockSelect struct {
	*ProductStockQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pss *ProductStockSelect) Aggregate(fns ...AggregateFunc) *ProductStockSelect {
	pss.fns = append(pss.fns, fns...)
	return pss
}

// Scan applies the selector query and scans the result into the given value.
func (pss *ProductStockSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pss.ctx, "Select")
	if err := pss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductStockQuery, *ProductStockSelect](ctx, pss.ProductStockQuery, pss, pss.inters, v)
}

func (pss *ProductStockSelect) sqlScan(ctx context.Context, root *ProductStockQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pss.fns))
	for _, fn := range pss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}