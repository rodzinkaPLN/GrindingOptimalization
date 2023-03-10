// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/parameter"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/predicate"
)

// ParameterDelete is the builder for deleting a Parameter entity.
type ParameterDelete struct {
	config
	hooks    []Hook
	mutation *ParameterMutation
}

// Where appends a list predicates to the ParameterDelete builder.
func (pd *ParameterDelete) Where(ps ...predicate.Parameter) *ParameterDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *ParameterDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ParameterMutation](ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *ParameterDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *ParameterDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: parameter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: parameter.FieldID,
			},
		},
	}
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// ParameterDeleteOne is the builder for deleting a single Parameter entity.
type ParameterDeleteOne struct {
	pd *ParameterDelete
}

// Where appends a list predicates to the ParameterDelete builder.
func (pdo *ParameterDeleteOne) Where(ps ...predicate.Parameter) *ParameterDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *ParameterDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{parameter.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *ParameterDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}
