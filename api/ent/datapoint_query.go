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
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/datapoint"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/dataset"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/predicate"
)

// DatapointQuery is the builder for querying Datapoint entities.
type DatapointQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.Datapoint
	withDatasets *DatasetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DatapointQuery builder.
func (dq *DatapointQuery) Where(ps ...predicate.Datapoint) *DatapointQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DatapointQuery) Limit(limit int) *DatapointQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DatapointQuery) Offset(offset int) *DatapointQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DatapointQuery) Unique(unique bool) *DatapointQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DatapointQuery) Order(o ...OrderFunc) *DatapointQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryDatasets chains the current query on the "datasets" edge.
func (dq *DatapointQuery) QueryDatasets() *DatasetQuery {
	query := (&DatasetClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(datapoint.Table, datapoint.FieldID, selector),
			sqlgraph.To(dataset.Table, dataset.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, datapoint.DatasetsTable, datapoint.DatasetsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Datapoint entity from the query.
// Returns a *NotFoundError when no Datapoint was found.
func (dq *DatapointQuery) First(ctx context.Context) (*Datapoint, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{datapoint.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DatapointQuery) FirstX(ctx context.Context) *Datapoint {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Datapoint ID from the query.
// Returns a *NotFoundError when no Datapoint ID was found.
func (dq *DatapointQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{datapoint.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DatapointQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Datapoint entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Datapoint entity is found.
// Returns a *NotFoundError when no Datapoint entities are found.
func (dq *DatapointQuery) Only(ctx context.Context) (*Datapoint, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{datapoint.Label}
	default:
		return nil, &NotSingularError{datapoint.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DatapointQuery) OnlyX(ctx context.Context) *Datapoint {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Datapoint ID in the query.
// Returns a *NotSingularError when more than one Datapoint ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DatapointQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{datapoint.Label}
	default:
		err = &NotSingularError{datapoint.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DatapointQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Datapoints.
func (dq *DatapointQuery) All(ctx context.Context) ([]*Datapoint, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Datapoint, *DatapointQuery]()
	return withInterceptors[[]*Datapoint](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DatapointQuery) AllX(ctx context.Context) []*Datapoint {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Datapoint IDs.
func (dq *DatapointQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err := dq.Select(datapoint.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DatapointQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DatapointQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DatapointQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DatapointQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DatapointQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DatapointQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DatapointQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DatapointQuery) Clone() *DatapointQuery {
	if dq == nil {
		return nil
	}
	return &DatapointQuery{
		config:       dq.config,
		ctx:          dq.ctx.Clone(),
		order:        append([]OrderFunc{}, dq.order...),
		inters:       append([]Interceptor{}, dq.inters...),
		predicates:   append([]predicate.Datapoint{}, dq.predicates...),
		withDatasets: dq.withDatasets.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithDatasets tells the query-builder to eager-load the nodes that are connected to
// the "datasets" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DatapointQuery) WithDatasets(opts ...func(*DatasetQuery)) *DatapointQuery {
	query := (&DatasetClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withDatasets = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DataTime time.Time `json:"data_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Datapoint.Query().
//		GroupBy(datapoint.FieldDataTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DatapointQuery) GroupBy(field string, fields ...string) *DatapointGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DatapointGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = datapoint.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DataTime time.Time `json:"data_time,omitempty"`
//	}
//
//	client.Datapoint.Query().
//		Select(datapoint.FieldDataTime).
//		Scan(ctx, &v)
func (dq *DatapointQuery) Select(fields ...string) *DatapointSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DatapointSelect{DatapointQuery: dq}
	sbuild.label = datapoint.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DatapointSelect configured with the given aggregations.
func (dq *DatapointQuery) Aggregate(fns ...AggregateFunc) *DatapointSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DatapointQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !datapoint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DatapointQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Datapoint, error) {
	var (
		nodes       = []*Datapoint{}
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withDatasets != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Datapoint).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Datapoint{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withDatasets; query != nil {
		if err := dq.loadDatasets(ctx, query, nodes, nil,
			func(n *Datapoint, e *Dataset) { n.Edges.Datasets = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DatapointQuery) loadDatasets(ctx context.Context, query *DatasetQuery, nodes []*Datapoint, init func(*Datapoint), assign func(*Datapoint, *Dataset)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Datapoint)
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

func (dq *DatapointQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DatapointQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   datapoint.Table,
			Columns: datapoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: datapoint.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, datapoint.FieldID)
		for i := range fields {
			if fields[i] != datapoint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DatapointQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(datapoint.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = datapoint.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DatapointGroupBy is the group-by builder for Datapoint entities.
type DatapointGroupBy struct {
	selector
	build *DatapointQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DatapointGroupBy) Aggregate(fns ...AggregateFunc) *DatapointGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DatapointGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DatapointQuery, *DatapointGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DatapointGroupBy) sqlScan(ctx context.Context, root *DatapointQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DatapointSelect is the builder for selecting fields of Datapoint entities.
type DatapointSelect struct {
	*DatapointQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DatapointSelect) Aggregate(fns ...AggregateFunc) *DatapointSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DatapointSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DatapointQuery, *DatapointSelect](ctx, ds.DatapointQuery, ds, ds.inters, v)
}

func (ds *DatapointSelect) sqlScan(ctx context.Context, root *DatapointQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
