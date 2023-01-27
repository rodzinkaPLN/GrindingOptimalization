// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/dataset"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/parameter"
)

// ParameterCreate is the builder for creating a Parameter entity.
type ParameterCreate struct {
	config
	mutation *ParameterMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ParameterCreate) SetName(s string) *ParameterCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetUnit sets the "unit" field.
func (pc *ParameterCreate) SetUnit(s string) *ParameterCreate {
	pc.mutation.SetUnit(s)
	return pc
}

// SetDatasetID sets the "dataset_id" field.
func (pc *ParameterCreate) SetDatasetID(u uuid.UUID) *ParameterCreate {
	pc.mutation.SetDatasetID(u)
	return pc
}

// SetNillableDatasetID sets the "dataset_id" field if the given value is not nil.
func (pc *ParameterCreate) SetNillableDatasetID(u *uuid.UUID) *ParameterCreate {
	if u != nil {
		pc.SetDatasetID(*u)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ParameterCreate) SetCreatedAt(t time.Time) *ParameterCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ParameterCreate) SetNillableCreatedAt(t *time.Time) *ParameterCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ParameterCreate) SetID(u uuid.UUID) *ParameterCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *ParameterCreate) SetNillableID(u *uuid.UUID) *ParameterCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetDatasetsID sets the "datasets" edge to the Dataset entity by ID.
func (pc *ParameterCreate) SetDatasetsID(id uuid.UUID) *ParameterCreate {
	pc.mutation.SetDatasetsID(id)
	return pc
}

// SetNillableDatasetsID sets the "datasets" edge to the Dataset entity by ID if the given value is not nil.
func (pc *ParameterCreate) SetNillableDatasetsID(id *uuid.UUID) *ParameterCreate {
	if id != nil {
		pc = pc.SetDatasetsID(*id)
	}
	return pc
}

// SetDatasets sets the "datasets" edge to the Dataset entity.
func (pc *ParameterCreate) SetDatasets(d *Dataset) *ParameterCreate {
	return pc.SetDatasetsID(d.ID)
}

// Mutation returns the ParameterMutation object of the builder.
func (pc *ParameterCreate) Mutation() *ParameterMutation {
	return pc.mutation
}

// Save creates the Parameter in the database.
func (pc *ParameterCreate) Save(ctx context.Context) (*Parameter, error) {
	pc.defaults()
	return withHooks[*Parameter, ParameterMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ParameterCreate) SaveX(ctx context.Context) *Parameter {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ParameterCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ParameterCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ParameterCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := parameter.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := parameter.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ParameterCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Parameter.name"`)}
	}
	if _, ok := pc.mutation.Unit(); !ok {
		return &ValidationError{Name: "unit", err: errors.New(`ent: missing required field "Parameter.unit"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Parameter.created_at"`)}
	}
	return nil
}

func (pc *ParameterCreate) sqlSave(ctx context.Context) (*Parameter, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ParameterCreate) createSpec() (*Parameter, *sqlgraph.CreateSpec) {
	var (
		_node = &Parameter{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: parameter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: parameter.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(parameter.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Unit(); ok {
		_spec.SetField(parameter.FieldUnit, field.TypeString, value)
		_node.Unit = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(parameter.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := pc.mutation.DatasetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parameter.DatasetsTable,
			Columns: []string{parameter.DatasetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dataset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DatasetID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ParameterCreateBulk is the builder for creating many Parameter entities in bulk.
type ParameterCreateBulk struct {
	config
	builders []*ParameterCreate
}

// Save creates the Parameter entities in the database.
func (pcb *ParameterCreateBulk) Save(ctx context.Context) ([]*Parameter, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Parameter, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ParameterMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ParameterCreateBulk) SaveX(ctx context.Context) []*Parameter {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ParameterCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ParameterCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
