// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/userinput"
)

// UserinputCreate is the builder for creating a Userinput entity.
type UserinputCreate struct {
	config
	mutation *UserinputMutation
	hooks    []Hook
}

// SetMin sets the "min" field.
func (uc *UserinputCreate) SetMin(f float64) *UserinputCreate {
	uc.mutation.SetMin(f)
	return uc
}

// SetMax sets the "max" field.
func (uc *UserinputCreate) SetMax(f float64) *UserinputCreate {
	uc.mutation.SetMax(f)
	return uc
}

// SetStep sets the "step" field.
func (uc *UserinputCreate) SetStep(f float64) *UserinputCreate {
	uc.mutation.SetStep(f)
	return uc
}

// SetDefaultval sets the "defaultval" field.
func (uc *UserinputCreate) SetDefaultval(f float64) *UserinputCreate {
	uc.mutation.SetDefaultval(f)
	return uc
}

// SetName sets the "name" field.
func (uc *UserinputCreate) SetName(s string) *UserinputCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetID sets the "id" field.
func (uc *UserinputCreate) SetID(u uuid.UUID) *UserinputCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserinputCreate) SetNillableID(u *uuid.UUID) *UserinputCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// Mutation returns the UserinputMutation object of the builder.
func (uc *UserinputCreate) Mutation() *UserinputMutation {
	return uc.mutation
}

// Save creates the Userinput in the database.
func (uc *UserinputCreate) Save(ctx context.Context) (*Userinput, error) {
	uc.defaults()
	return withHooks[*Userinput, UserinputMutation](ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserinputCreate) SaveX(ctx context.Context) *Userinput {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserinputCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserinputCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserinputCreate) defaults() {
	if _, ok := uc.mutation.ID(); !ok {
		v := userinput.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserinputCreate) check() error {
	if _, ok := uc.mutation.Min(); !ok {
		return &ValidationError{Name: "min", err: errors.New(`ent: missing required field "Userinput.min"`)}
	}
	if _, ok := uc.mutation.Max(); !ok {
		return &ValidationError{Name: "max", err: errors.New(`ent: missing required field "Userinput.max"`)}
	}
	if _, ok := uc.mutation.Step(); !ok {
		return &ValidationError{Name: "step", err: errors.New(`ent: missing required field "Userinput.step"`)}
	}
	if _, ok := uc.mutation.Defaultval(); !ok {
		return &ValidationError{Name: "defaultval", err: errors.New(`ent: missing required field "Userinput.defaultval"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Userinput.name"`)}
	}
	return nil
}

func (uc *UserinputCreate) sqlSave(ctx context.Context) (*Userinput, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserinputCreate) createSpec() (*Userinput, *sqlgraph.CreateSpec) {
	var (
		_node = &Userinput{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userinput.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userinput.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.Min(); ok {
		_spec.SetField(userinput.FieldMin, field.TypeFloat64, value)
		_node.Min = value
	}
	if value, ok := uc.mutation.Max(); ok {
		_spec.SetField(userinput.FieldMax, field.TypeFloat64, value)
		_node.Max = value
	}
	if value, ok := uc.mutation.Step(); ok {
		_spec.SetField(userinput.FieldStep, field.TypeFloat64, value)
		_node.Step = value
	}
	if value, ok := uc.mutation.Defaultval(); ok {
		_spec.SetField(userinput.FieldDefaultval, field.TypeFloat64, value)
		_node.Defaultval = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(userinput.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// UserinputCreateBulk is the builder for creating many Userinput entities in bulk.
type UserinputCreateBulk struct {
	config
	builders []*UserinputCreate
}

// Save creates the Userinput entities in the database.
func (ucb *UserinputCreateBulk) Save(ctx context.Context) ([]*Userinput, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Userinput, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserinputMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserinputCreateBulk) SaveX(ctx context.Context) []*Userinput {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserinputCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserinputCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
