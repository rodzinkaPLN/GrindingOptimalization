// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/predicate"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/userinput"
)

// UserinputDelete is the builder for deleting a Userinput entity.
type UserinputDelete struct {
	config
	hooks    []Hook
	mutation *UserinputMutation
}

// Where appends a list predicates to the UserinputDelete builder.
func (ud *UserinputDelete) Where(ps ...predicate.Userinput) *UserinputDelete {
	ud.mutation.Where(ps...)
	return ud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ud *UserinputDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, UserinputMutation](ctx, ud.sqlExec, ud.mutation, ud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ud *UserinputDelete) ExecX(ctx context.Context) int {
	n, err := ud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ud *UserinputDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userinput.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userinput.FieldID,
			},
		},
	}
	if ps := ud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ud.mutation.done = true
	return affected, err
}

// UserinputDeleteOne is the builder for deleting a single Userinput entity.
type UserinputDeleteOne struct {
	ud *UserinputDelete
}

// Where appends a list predicates to the UserinputDelete builder.
func (udo *UserinputDeleteOne) Where(ps ...predicate.Userinput) *UserinputDeleteOne {
	udo.ud.mutation.Where(ps...)
	return udo
}

// Exec executes the deletion query.
func (udo *UserinputDeleteOne) Exec(ctx context.Context) error {
	n, err := udo.ud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userinput.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (udo *UserinputDeleteOne) ExecX(ctx context.Context) {
	if err := udo.Exec(ctx); err != nil {
		panic(err)
	}
}