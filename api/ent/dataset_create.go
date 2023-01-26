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

// DatasetCreate is the builder for creating a Dataset entity.
type DatasetCreate struct {
	config
	mutation *DatasetMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dc *DatasetCreate) SetName(s string) *DatasetCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetCreatedAt sets the "created_at" field.
func (dc *DatasetCreate) SetCreatedAt(t time.Time) *DatasetCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DatasetCreate) SetNillableCreatedAt(t *time.Time) *DatasetCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DatasetCreate) SetID(u uuid.UUID) *DatasetCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DatasetCreate) SetNillableID(u *uuid.UUID) *DatasetCreate {
	if u != nil {
		dc.SetID(*u)
	}
	return dc
}

// AddParameterIDs adds the "parameters" edge to the Parameter entity by IDs.
func (dc *DatasetCreate) AddParameterIDs(ids ...uuid.UUID) *DatasetCreate {
	dc.mutation.AddParameterIDs(ids...)
	return dc
}

// AddParameters adds the "parameters" edges to the Parameter entity.
func (dc *DatasetCreate) AddParameters(p ...*Parameter) *DatasetCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return dc.AddParameterIDs(ids...)
}

// Mutation returns the DatasetMutation object of the builder.
func (dc *DatasetCreate) Mutation() *DatasetMutation {
	return dc.mutation
}

// Save creates the Dataset in the database.
func (dc *DatasetCreate) Save(ctx context.Context) (*Dataset, error) {
	dc.defaults()
	return withHooks[*Dataset, DatasetMutation](ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DatasetCreate) SaveX(ctx context.Context) *Dataset {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DatasetCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DatasetCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DatasetCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := dataset.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := dataset.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DatasetCreate) check() error {
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Dataset.name"`)}
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Dataset.created_at"`)}
	}
	return nil
}

func (dc *DatasetCreate) sqlSave(ctx context.Context) (*Dataset, error) {
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

func (dc *DatasetCreate) createSpec() (*Dataset, *sqlgraph.CreateSpec) {
	var (
		_node = &Dataset{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dataset.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: dataset.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(dataset.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(dataset.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := dc.mutation.ParametersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dataset.ParametersTable,
			Columns: []string{dataset.ParametersColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DatasetCreateBulk is the builder for creating many Dataset entities in bulk.
type DatasetCreateBulk struct {
	config
	builders []*DatasetCreate
}

// Save creates the Dataset entities in the database.
func (dcb *DatasetCreateBulk) Save(ctx context.Context) ([]*Dataset, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dataset, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DatasetMutation)
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
func (dcb *DatasetCreateBulk) SaveX(ctx context.Context) []*Dataset {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DatasetCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DatasetCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}