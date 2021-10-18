// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dictor/jaeminbot/ent/command"
	"github.com/dictor/jaeminbot/ent/predicate"
	"github.com/dictor/jaeminbot/ent/resultlog"
)

// ResultLogQuery is the builder for querying ResultLog entities.
type ResultLogQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ResultLog
	// eager-loading edges.
	withCommand *CommandQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ResultLogQuery builder.
func (rlq *ResultLogQuery) Where(ps ...predicate.ResultLog) *ResultLogQuery {
	rlq.predicates = append(rlq.predicates, ps...)
	return rlq
}

// Limit adds a limit step to the query.
func (rlq *ResultLogQuery) Limit(limit int) *ResultLogQuery {
	rlq.limit = &limit
	return rlq
}

// Offset adds an offset step to the query.
func (rlq *ResultLogQuery) Offset(offset int) *ResultLogQuery {
	rlq.offset = &offset
	return rlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rlq *ResultLogQuery) Unique(unique bool) *ResultLogQuery {
	rlq.unique = &unique
	return rlq
}

// Order adds an order step to the query.
func (rlq *ResultLogQuery) Order(o ...OrderFunc) *ResultLogQuery {
	rlq.order = append(rlq.order, o...)
	return rlq
}

// QueryCommand chains the current query on the "command" edge.
func (rlq *ResultLogQuery) QueryCommand() *CommandQuery {
	query := &CommandQuery{config: rlq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resultlog.Table, resultlog.FieldID, selector),
			sqlgraph.To(command.Table, command.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, resultlog.CommandTable, resultlog.CommandPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ResultLog entity from the query.
// Returns a *NotFoundError when no ResultLog was found.
func (rlq *ResultLogQuery) First(ctx context.Context) (*ResultLog, error) {
	nodes, err := rlq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{resultlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rlq *ResultLogQuery) FirstX(ctx context.Context) *ResultLog {
	node, err := rlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ResultLog ID from the query.
// Returns a *NotFoundError when no ResultLog ID was found.
func (rlq *ResultLogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rlq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{resultlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rlq *ResultLogQuery) FirstIDX(ctx context.Context) int {
	id, err := rlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ResultLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ResultLog entity is not found.
// Returns a *NotFoundError when no ResultLog entities are found.
func (rlq *ResultLogQuery) Only(ctx context.Context) (*ResultLog, error) {
	nodes, err := rlq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{resultlog.Label}
	default:
		return nil, &NotSingularError{resultlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rlq *ResultLogQuery) OnlyX(ctx context.Context) *ResultLog {
	node, err := rlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ResultLog ID in the query.
// Returns a *NotSingularError when exactly one ResultLog ID is not found.
// Returns a *NotFoundError when no entities are found.
func (rlq *ResultLogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rlq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{resultlog.Label}
	default:
		err = &NotSingularError{resultlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rlq *ResultLogQuery) OnlyIDX(ctx context.Context) int {
	id, err := rlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ResultLogs.
func (rlq *ResultLogQuery) All(ctx context.Context) ([]*ResultLog, error) {
	if err := rlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rlq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rlq *ResultLogQuery) AllX(ctx context.Context) []*ResultLog {
	nodes, err := rlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ResultLog IDs.
func (rlq *ResultLogQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rlq.Select(resultlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rlq *ResultLogQuery) IDsX(ctx context.Context) []int {
	ids, err := rlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rlq *ResultLogQuery) Count(ctx context.Context) (int, error) {
	if err := rlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rlq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rlq *ResultLogQuery) CountX(ctx context.Context) int {
	count, err := rlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rlq *ResultLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := rlq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rlq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rlq *ResultLogQuery) ExistX(ctx context.Context) bool {
	exist, err := rlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ResultLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rlq *ResultLogQuery) Clone() *ResultLogQuery {
	if rlq == nil {
		return nil
	}
	return &ResultLogQuery{
		config:      rlq.config,
		limit:       rlq.limit,
		offset:      rlq.offset,
		order:       append([]OrderFunc{}, rlq.order...),
		predicates:  append([]predicate.ResultLog{}, rlq.predicates...),
		withCommand: rlq.withCommand.Clone(),
		// clone intermediate query.
		sql:  rlq.sql.Clone(),
		path: rlq.path,
	}
}

// WithCommand tells the query-builder to eager-load the nodes that are connected to
// the "command" edge. The optional arguments are used to configure the query builder of the edge.
func (rlq *ResultLogQuery) WithCommand(opts ...func(*CommandQuery)) *ResultLogQuery {
	query := &CommandQuery{config: rlq.config}
	for _, opt := range opts {
		opt(query)
	}
	rlq.withCommand = query
	return rlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Level string `json:"level,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ResultLog.Query().
//		GroupBy(resultlog.FieldLevel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rlq *ResultLogQuery) GroupBy(field string, fields ...string) *ResultLogGroupBy {
	group := &ResultLogGroupBy{config: rlq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rlq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Level string `json:"level,omitempty"`
//	}
//
//	client.ResultLog.Query().
//		Select(resultlog.FieldLevel).
//		Scan(ctx, &v)
//
func (rlq *ResultLogQuery) Select(fields ...string) *ResultLogSelect {
	rlq.fields = append(rlq.fields, fields...)
	return &ResultLogSelect{ResultLogQuery: rlq}
}

func (rlq *ResultLogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rlq.fields {
		if !resultlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rlq.path != nil {
		prev, err := rlq.path(ctx)
		if err != nil {
			return err
		}
		rlq.sql = prev
	}
	return nil
}

func (rlq *ResultLogQuery) sqlAll(ctx context.Context) ([]*ResultLog, error) {
	var (
		nodes       = []*ResultLog{}
		_spec       = rlq.querySpec()
		loadedTypes = [1]bool{
			rlq.withCommand != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ResultLog{config: rlq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, rlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rlq.withCommand; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*ResultLog, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Command = []*Command{}
		}
		var (
			edgeids []string
			edges   = make(map[string][]*ResultLog)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   resultlog.CommandTable,
				Columns: resultlog.CommandPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(resultlog.CommandPrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullString)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullString)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := ein.String
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, rlq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "command": %w`, err)
		}
		query.Where(command.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "command" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Command = append(nodes[i].Edges.Command, n)
			}
		}
	}

	return nodes, nil
}

func (rlq *ResultLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rlq.querySpec()
	return sqlgraph.CountNodes(ctx, rlq.driver, _spec)
}

func (rlq *ResultLogQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rlq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (rlq *ResultLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resultlog.Table,
			Columns: resultlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resultlog.FieldID,
			},
		},
		From:   rlq.sql,
		Unique: true,
	}
	if unique := rlq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rlq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resultlog.FieldID)
		for i := range fields {
			if fields[i] != resultlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rlq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rlq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rlq *ResultLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rlq.driver.Dialect())
	t1 := builder.Table(resultlog.Table)
	columns := rlq.fields
	if len(columns) == 0 {
		columns = resultlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rlq.sql != nil {
		selector = rlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range rlq.predicates {
		p(selector)
	}
	for _, p := range rlq.order {
		p(selector)
	}
	if offset := rlq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rlq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ResultLogGroupBy is the group-by builder for ResultLog entities.
type ResultLogGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rlgb *ResultLogGroupBy) Aggregate(fns ...AggregateFunc) *ResultLogGroupBy {
	rlgb.fns = append(rlgb.fns, fns...)
	return rlgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rlgb *ResultLogGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rlgb.path(ctx)
	if err != nil {
		return err
	}
	rlgb.sql = query
	return rlgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rlgb *ResultLogGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rlgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (rlgb *ResultLogGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rlgb.fields) > 1 {
		return nil, errors.New("ent: ResultLogGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rlgb *ResultLogGroupBy) StringsX(ctx context.Context) []string {
	v, err := rlgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rlgb *ResultLogGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rlgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resultlog.Label}
	default:
		err = fmt.Errorf("ent: ResultLogGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rlgb *ResultLogGroupBy) StringX(ctx context.Context) string {
	v, err := rlgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (rlgb *ResultLogGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rlgb.fields) > 1 {
		return nil, errors.New("ent: ResultLogGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rlgb *ResultLogGroupBy) IntsX(ctx context.Context) []int {
	v, err := rlgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rlgb *ResultLogGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rlgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resultlog.Label}
	default:
		err = fmt.Errorf("ent: ResultLogGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rlgb *ResultLogGroupBy) IntX(ctx context.Context) int {
	v, err := rlgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (rlgb *ResultLogGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rlgb.fields) > 1 {
		return nil, errors.New("ent: ResultLogGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rlgb *ResultLogGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rlgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rlgb *ResultLogGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rlgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resultlog.Label}
	default:
		err = fmt.Errorf("ent: ResultLogGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rlgb *ResultLogGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rlgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (rlgb *ResultLogGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rlgb.fields) > 1 {
		return nil, errors.New("ent: ResultLogGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rlgb *ResultLogGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rlgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rlgb *ResultLogGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rlgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resultlog.Label}
	default:
		err = fmt.Errorf("ent: ResultLogGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rlgb *ResultLogGroupBy) BoolX(ctx context.Context) bool {
	v, err := rlgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rlgb *ResultLogGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rlgb.fields {
		if !resultlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rlgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rlgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rlgb *ResultLogGroupBy) sqlQuery() *sql.Selector {
	selector := rlgb.sql.Select()
	aggregation := make([]string, 0, len(rlgb.fns))
	for _, fn := range rlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rlgb.fields)+len(rlgb.fns))
		for _, f := range rlgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rlgb.fields...)...)
}

// ResultLogSelect is the builder for selecting fields of ResultLog entities.
type ResultLogSelect struct {
	*ResultLogQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rls *ResultLogSelect) Scan(ctx context.Context, v interface{}) error {
	if err := rls.prepareQuery(ctx); err != nil {
		return err
	}
	rls.sql = rls.ResultLogQuery.sqlQuery(ctx)
	return rls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rls *ResultLogSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (rls *ResultLogSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rls.fields) > 1 {
		return nil, errors.New("ent: ResultLogSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rls *ResultLogSelect) StringsX(ctx context.Context) []string {
	v, err := rls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (rls *ResultLogSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resultlog.Label}
	default:
		err = fmt.Errorf("ent: ResultLogSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rls *ResultLogSelect) StringX(ctx context.Context) string {
	v, err := rls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (rls *ResultLogSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rls.fields) > 1 {
		return nil, errors.New("ent: ResultLogSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rls *ResultLogSelect) IntsX(ctx context.Context) []int {
	v, err := rls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (rls *ResultLogSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resultlog.Label}
	default:
		err = fmt.Errorf("ent: ResultLogSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rls *ResultLogSelect) IntX(ctx context.Context) int {
	v, err := rls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (rls *ResultLogSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rls.fields) > 1 {
		return nil, errors.New("ent: ResultLogSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rls *ResultLogSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (rls *ResultLogSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resultlog.Label}
	default:
		err = fmt.Errorf("ent: ResultLogSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rls *ResultLogSelect) Float64X(ctx context.Context) float64 {
	v, err := rls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (rls *ResultLogSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rls.fields) > 1 {
		return nil, errors.New("ent: ResultLogSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rls *ResultLogSelect) BoolsX(ctx context.Context) []bool {
	v, err := rls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (rls *ResultLogSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resultlog.Label}
	default:
		err = fmt.Errorf("ent: ResultLogSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rls *ResultLogSelect) BoolX(ctx context.Context) bool {
	v, err := rls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rls *ResultLogSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rls.sql.Query()
	if err := rls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
