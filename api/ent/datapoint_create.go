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
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/datapoint"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/parameter"
)

// DatapointCreate is the builder for creating a Datapoint entity.
type DatapointCreate struct {
	config
	mutation *DatapointMutation
	hooks    []Hook
}

// SetDataTime sets the "data_time" field.
func (dc *DatapointCreate) SetDataTime(t time.Time) *DatapointCreate {
	dc.mutation.SetDataTime(t)
	return dc
}

// SetParameterID sets the "parameter_id" field.
func (dc *DatapointCreate) SetParameterID(u uuid.UUID) *DatapointCreate {
	dc.mutation.SetParameterID(u)
	return dc
}

// SetNillableParameterID sets the "parameter_id" field if the given value is not nil.
func (dc *DatapointCreate) SetNillableParameterID(u *uuid.UUID) *DatapointCreate {
	if u != nil {
		dc.SetParameterID(*u)
	}
	return dc
}

// SetVal sets the "val" field.
func (dc *DatapointCreate) SetVal(f float64) *DatapointCreate {
	dc.mutation.SetVal(f)
	return dc
}

// SetCreatedAt sets the "created_at" field.
func (dc *DatapointCreate) SetCreatedAt(t time.Time) *DatapointCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DatapointCreate) SetNillableCreatedAt(t *time.Time) *DatapointCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DatapointCreate) SetID(u uuid.UUID) *DatapointCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DatapointCreate) SetNillableID(u *uuid.UUID) *DatapointCreate {
	if u != nil {
		dc.SetID(*u)
	}
	return dc
}

// SetParametersID sets the "parameters" edge to the Parameter entity by ID.
func (dc *DatapointCreate) SetParametersID(id uuid.UUID) *DatapointCreate {
	dc.mutation.SetParametersID(id)
	return dc
}

// SetNillableParametersID sets the "parameters" edge to the Parameter entity by ID if the given value is not nil.
func (dc *DatapointCreate) SetNillableParametersID(id *uuid.UUID) *DatapointCreate {
	if id != nil {
		dc = dc.SetParametersID(*id)
	}
	return dc
}

// SetParameters sets the "parameters" edge to the Parameter entity.
func (dc *DatapointCreate) SetParameters(p *Parameter) *DatapointCreate {
	return dc.SetParametersID(p.ID)
}

// Mutation returns the DatapointMutation object of the builder.
func (dc *DatapointCreate) Mutation() *DatapointMutation {
	return dc.mutation
}

// Save creates the Datapoint in the database.
func (dc *DatapointCreate) Save(ctx context.Context) (*Datapoint, error) {
	dc.defaults()
	return withHooks[*Datapoint, DatapointMutation](ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DatapointCreate) SaveX(ctx context.Context) *Datapoint {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DatapointCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DatapointCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DatapointCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := datapoint.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := datapoint.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DatapointCreate) check() error {
	if _, ok := dc.mutation.DataTime(); !ok {
		return &ValidationError{Name: "data_time", err: errors.New(`ent: missing required field "Datapoint.data_time"`)}
	}
	if _, ok := dc.mutation.Val(); !ok {
		return &ValidationError{Name: "val", err: errors.New(`ent: missing required field "Datapoint.val"`)}
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Datapoint.created_at"`)}
	}
	return nil
}

func (dc *DatapointCreate) sqlSave(ctx context.Context) (*Datapoint, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
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
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DatapointCreate) createSpec() (*Datapoint, *sqlgraph.CreateSpec) {
	var (
		_node = &Datapoint{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: datapoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: datapoint.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.DataTime(); ok {
		_spec.SetField(datapoint.FieldDataTime, field.TypeTime, value)
		_node.DataTime = value
	}
	if value, ok := dc.mutation.Val(); ok {
		_spec.SetField(datapoint.FieldVal, field.TypeFloat64, value)
		_node.Val = value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(datapoint.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := dc.mutation.ParametersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   datapoint.ParametersTable,
			Columns: []string{datapoint.ParametersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: parameter.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParameterID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DatapointCreateBulk is the builder for creating many Datapoint entities in bulk.
type DatapointCreateBulk struct {
	config
	builders []*DatapointCreate
}

// Save creates the Datapoint entities in the database.
func (dcb *DatapointCreateBulk) Save(ctx context.Context) ([]*Datapoint, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Datapoint, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DatapointMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DatapointCreateBulk) SaveX(ctx context.Context) []*Datapoint {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DatapointCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DatapointCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
