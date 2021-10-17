// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dictor/jaeminbot/ent/predicate"
	"github.com/dictor/jaeminbot/ent/resultlog"
)

// ResultLogDelete is the builder for deleting a ResultLog entity.
type ResultLogDelete struct {
	config
	hooks    []Hook
	mutation *ResultLogMutation
}

// Where appends a list predicates to the ResultLogDelete builder.
func (rld *ResultLogDelete) Where(ps ...predicate.ResultLog) *ResultLogDelete {
	rld.mutation.Where(ps...)
	return rld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rld *ResultLogDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rld.hooks) == 0 {
		affected, err = rld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResultLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rld.mutation = mutation
			affected, err = rld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rld.hooks) - 1; i >= 0; i-- {
			if rld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rld *ResultLogDelete) ExecX(ctx context.Context) int {
	n, err := rld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rld *ResultLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: resultlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resultlog.FieldID,
			},
		},
	}
	if ps := rld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rld.driver, _spec)
}

// ResultLogDeleteOne is the builder for deleting a single ResultLog entity.
type ResultLogDeleteOne struct {
	rld *ResultLogDelete
}

// Exec executes the deletion query.
func (rldo *ResultLogDeleteOne) Exec(ctx context.Context) error {
	n, err := rldo.rld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{resultlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rldo *ResultLogDeleteOne) ExecX(ctx context.Context) {
	rldo.rld.ExecX(ctx)
}
