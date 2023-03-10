// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/dataset"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/parameter"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/predicate"
)

// ParameterQuery is the builder for querying Parameter entities.
type ParameterQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.Parameter
	withDatasets *DatasetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ParameterQuery builder.
func (pq *ParameterQuery) Where(ps ...predicate.Parameter) *ParameterQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *ParameterQuery) Limit(limit int) *ParameterQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *ParameterQuery) Offset(offset int) *ParameterQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ParameterQuery) Unique(unique bool) *ParameterQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *ParameterQuery) Order(o ...OrderFunc) *ParameterQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryDatasets chains the current query on the "datasets" edge.
func (pq *ParameterQuery) QueryDatasets() *DatasetQuery {
	query := (&DatasetClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(parameter.Table, parameter.FieldID, selector),
			sqlgraph.To(dataset.Table, dataset.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, parameter.DatasetsTable, parameter.DatasetsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Parameter entity from the query.
// Returns a *NotFoundError when no Parameter was found.
func (pq *ParameterQuery) First(ctx context.Context) (*Parameter, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{parameter.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ParameterQuery) FirstX(ctx context.Context) *Parameter {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Parameter ID from the query.
// Returns a *NotFoundError when no Parameter ID was found.
func (pq *ParameterQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{parameter.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ParameterQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Parameter entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Parameter entity is found.
// Returns a *NotFoundError when no Parameter entities are found.
func (pq *ParameterQuery) Only(ctx context.Context) (*Parameter, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{parameter.Label}
	default:
		return nil, &NotSingularError{parameter.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ParameterQuery) OnlyX(ctx context.Context) *Parameter {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Parameter ID in the query.
// Returns a *NotSingularError when more than one Parameter ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ParameterQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{parameter.Label}
	default:
		err = &NotSingularError{parameter.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ParameterQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Parameters.
func (pq *ParameterQuery) All(ctx context.Context) ([]*Parameter, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Parameter, *ParameterQuery]()
	return withInterceptors[[]*Parameter](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *ParameterQuery) AllX(ctx context.Context) []*Parameter {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Parameter IDs.
func (pq *ParameterQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err := pq.Select(parameter.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ParameterQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ParameterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*ParameterQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ParameterQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ParameterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ParameterQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ParameterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ParameterQuery) Clone() *ParameterQuery {
	if pq == nil {
		return nil
	}
	return &ParameterQuery{
		config:       pq.config,
		ctx:          pq.ctx.Clone(),
		order:        append([]OrderFunc{}, pq.order...),
		inters:       append([]Interceptor{}, pq.inters...),
		predicates:   append([]predicate.Parameter{}, pq.predicates...),
		withDatasets: pq.withDatasets.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithDatasets tells the query-builder to eager-load the nodes that are connected to
// the "datasets" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ParameterQuery) WithDatasets(opts ...func(*DatasetQuery)) *ParameterQuery {
	query := (&DatasetClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withDatasets = query
	return pq
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
//	client.Parameter.Query().
//		GroupBy(parameter.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *ParameterQuery) GroupBy(field string, fields ...string) *ParameterGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ParameterGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = parameter.Label
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
//	client.Parameter.Query().
//		Select(parameter.FieldName).
//		Scan(ctx, &v)
func (pq *ParameterQuery) Select(fields ...string) *ParameterSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &ParameterSelect{ParameterQuery: pq}
	sbuild.label = parameter.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ParameterSelect configured with the given aggregations.
func (pq *ParameterQuery) Aggregate(fns ...AggregateFunc) *ParameterSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ParameterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !parameter.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *ParameterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Parameter, error) {
	var (
		nodes       = []*Parameter{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withDatasets != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Parameter).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Parameter{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withDatasets; query != nil {
		if err := pq.loadDatasets(ctx, query, nodes, nil,
			func(n *Parameter, e *Dataset) { n.Edges.Datasets = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *ParameterQuery) loadDatasets(ctx context.Context, query *DatasetQuery, nodes []*Parameter, init func(*Parameter), assign func(*Parameter, *Dataset)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Parameter)
	for i := range nodes {
		fk := nodes[i].DatasetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(dataset.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dataset_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *ParameterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ParameterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   parameter.Table,
			Columns: parameter.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: parameter.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, parameter.FieldID)
		for i := range fields {
			if fields[i] != parameter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ParameterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(parameter.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = parameter.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ParameterGroupBy is the group-by builder for Parameter entities.
type ParameterGroupBy struct {
	selector
	build *ParameterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ParameterGroupBy) Aggregate(fns ...AggregateFunc) *ParameterGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *ParameterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ParameterQuery, *ParameterGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *ParameterGroupBy) sqlScan(ctx context.Context, root *ParameterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ParameterSelect is the builder for selecting fields of Parameter entities.
type ParameterSelect struct {
	*ParameterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ParameterSelect) Aggregate(fns ...AggregateFunc) *ParameterSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ParameterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ParameterQuery, *ParameterSelect](ctx, ps.ParameterQuery, ps, ps.inters, v)
}

func (ps *ParameterSelect) sqlScan(ctx context.Context, root *ParameterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
