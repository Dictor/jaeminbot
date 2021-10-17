// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dictor/jaeminbot/ent/command"
	"github.com/dictor/jaeminbot/ent/resultlog"
)

// ResultLogCreate is the builder for creating a ResultLog entity.
type ResultLogCreate struct {
	config
	mutation *ResultLogMutation
	hooks    []Hook
}

// SetLevel sets the "level" field.
func (rlc *ResultLogCreate) SetLevel(s string) *ResultLogCreate {
	rlc.mutation.SetLevel(s)
	return rlc
}

// SetLog sets the "log" field.
func (rlc *ResultLogCreate) SetLog(s string) *ResultLogCreate {
	rlc.mutation.SetLog(s)
	return rlc
}

// AddCommandIDs adds the "command" edge to the Command entity by IDs.
func (rlc *ResultLogCreate) AddCommandIDs(ids ...int) *ResultLogCreate {
	rlc.mutation.AddCommandIDs(ids...)
	return rlc
}

// AddCommand adds the "command" edges to the Command entity.
func (rlc *ResultLogCreate) AddCommand(c ...*Command) *ResultLogCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return rlc.AddCommandIDs(ids...)
}

// Mutation returns the ResultLogMutation object of the builder.
func (rlc *ResultLogCreate) Mutation() *ResultLogMutation {
	return rlc.mutation
}

// Save creates the ResultLog in the database.
func (rlc *ResultLogCreate) Save(ctx context.Context) (*ResultLog, error) {
	var (
		err  error
		node *ResultLog
	)
	if len(rlc.hooks) == 0 {
		if err = rlc.check(); err != nil {
			return nil, err
		}
		node, err = rlc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResultLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rlc.check(); err != nil {
				return nil, err
			}
			rlc.mutation = mutation
			if node, err = rlc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rlc.hooks) - 1; i >= 0; i-- {
			if rlc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rlc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rlc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rlc *ResultLogCreate) SaveX(ctx context.Context) *ResultLog {
	v, err := rlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rlc *ResultLogCreate) Exec(ctx context.Context) error {
	_, err := rlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rlc *ResultLogCreate) ExecX(ctx context.Context) {
	if err := rlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rlc *ResultLogCreate) check() error {
	if _, ok := rlc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "level"`)}
	}
	if _, ok := rlc.mutation.Log(); !ok {
		return &ValidationError{Name: "log", err: errors.New(`ent: missing required field "log"`)}
	}
	return nil
}

func (rlc *ResultLogCreate) sqlSave(ctx context.Context) (*ResultLog, error) {
	_node, _spec := rlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rlc *ResultLogCreate) createSpec() (*ResultLog, *sqlgraph.CreateSpec) {
	var (
		_node = &ResultLog{config: rlc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: resultlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resultlog.FieldID,
			},
		}
	)
	if value, ok := rlc.mutation.Level(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resultlog.FieldLevel,
		})
		_node.Level = value
	}
	if value, ok := rlc.mutation.Log(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resultlog.FieldLog,
		})
		_node.Log = value
	}
	if nodes := rlc.mutation.CommandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   resultlog.CommandTable,
			Columns: resultlog.CommandPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: command.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ResultLogCreateBulk is the builder for creating many ResultLog entities in bulk.
type ResultLogCreateBulk struct {
	config
	builders []*ResultLogCreate
}

// Save creates the ResultLog entities in the database.
func (rlcb *ResultLogCreateBulk) Save(ctx context.Context) ([]*ResultLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rlcb.builders))
	nodes := make([]*ResultLog, len(rlcb.builders))
	mutators := make([]Mutator, len(rlcb.builders))
	for i := range rlcb.builders {
		func(i int, root context.Context) {
			builder := rlcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResultLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rlcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rlcb *ResultLogCreateBulk) SaveX(ctx context.Context) []*ResultLog {
	v, err := rlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rlcb *ResultLogCreateBulk) Exec(ctx context.Context) error {
	_, err := rlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rlcb *ResultLogCreateBulk) ExecX(ctx context.Context) {
	if err := rlcb.Exec(ctx); err != nil {
		panic(err)
	}
}
