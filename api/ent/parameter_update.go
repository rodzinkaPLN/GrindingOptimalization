// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/dataset"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/parameter"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/predicate"
)

// ParameterUpdate is the builder for updating Parameter entities.
type ParameterUpdate struct {
	config
	hooks    []Hook
	mutation *ParameterMutation
}

// Where appends a list predicates to the ParameterUpdate builder.
func (pu *ParameterUpdate) Where(ps ...predicate.Parameter) *ParameterUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ParameterUpdate) SetName(s string) *ParameterUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetUnit sets the "unit" field.
func (pu *ParameterUpdate) SetUnit(s string) *ParameterUpdate {
	pu.mutation.SetUnit(s)
	return pu
}

// SetDatasetID sets the "dataset_id" field.
func (pu *ParameterUpdate) SetDatasetID(u uuid.UUID) *ParameterUpdate {
	pu.mutation.SetDatasetID(u)
	return pu
}

// SetNillableDatasetID sets the "dataset_id" field if the given value is not nil.
func (pu *ParameterUpdate) SetNillableDatasetID(u *uuid.UUID) *ParameterUpdate {
	if u != nil {
		pu.SetDatasetID(*u)
	}
	return pu
}

// ClearDatasetID clears the value of the "dataset_id" field.
func (pu *ParameterUpdate) ClearDatasetID() *ParameterUpdate {
	pu.mutation.ClearDatasetID()
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *ParameterUpdate) SetCreatedAt(t time.Time) *ParameterUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *ParameterUpdate) SetNillableCreatedAt(t *time.Time) *ParameterUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetDatasetsID sets the "datasets" edge to the Dataset entity by ID.
func (pu *ParameterUpdate) SetDatasetsID(id uuid.UUID) *ParameterUpdate {
	pu.mutation.SetDatasetsID(id)
	return pu
}

// SetNillableDatasetsID sets the "datasets" edge to the Dataset entity by ID if the given value is not nil.
func (pu *ParameterUpdate) SetNillableDatasetsID(id *uuid.UUID) *ParameterUpdate {
	if id != nil {
		pu = pu.SetDatasetsID(*id)
	}
	return pu
}

// SetDatasets sets the "datasets" edge to the Dataset entity.
func (pu *ParameterUpdate) SetDatasets(d *Dataset) *ParameterUpdate {
	return pu.SetDatasetsID(d.ID)
}

// Mutation returns the ParameterMutation object of the builder.
func (pu *ParameterUpdate) Mutation() *ParameterMutation {
	return pu.mutation
}

// ClearDatasets clears the "datasets" edge to the Dataset entity.
func (pu *ParameterUpdate) ClearDatasets() *ParameterUpdate {
	pu.mutation.ClearDatasets()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ParameterUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ParameterMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ParameterUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ParameterUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ParameterUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ParameterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   parameter.Table,
			Columns: parameter.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: parameter.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(parameter.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Unit(); ok {
		_spec.SetField(parameter.FieldUnit, field.TypeString, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(parameter.FieldCreatedAt, field.TypeTime, value)
	}
	if pu.mutation.DatasetsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DatasetsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{parameter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ParameterUpdateOne is the builder for updating a single Parameter entity.
type ParameterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ParameterMutation
}

// SetName sets the "name" field.
func (puo *ParameterUpdateOne) SetName(s string) *ParameterUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetUnit sets the "unit" field.
func (puo *ParameterUpdateOne) SetUnit(s string) *ParameterUpdateOne {
	puo.mutation.SetUnit(s)
	return puo
}

// SetDatasetID sets the "dataset_id" field.
func (puo *ParameterUpdateOne) SetDatasetID(u uuid.UUID) *ParameterUpdateOne {
	puo.mutation.SetDatasetID(u)
	return puo
}

// SetNillableDatasetID sets the "dataset_id" field if the given value is not nil.
func (puo *ParameterUpdateOne) SetNillableDatasetID(u *uuid.UUID) *ParameterUpdateOne {
	if u != nil {
		puo.SetDatasetID(*u)
	}
	return puo
}

// ClearDatasetID clears the value of the "dataset_id" field.
func (puo *ParameterUpdateOne) ClearDatasetID() *ParameterUpdateOne {
	puo.mutation.ClearDatasetID()
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *ParameterUpdateOne) SetCreatedAt(t time.Time) *ParameterUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *ParameterUpdateOne) SetNillableCreatedAt(t *time.Time) *ParameterUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetDatasetsID sets the "datasets" edge to the Dataset entity by ID.
func (puo *ParameterUpdateOne) SetDatasetsID(id uuid.UUID) *ParameterUpdateOne {
	puo.mutation.SetDatasetsID(id)
	return puo
}

// SetNillableDatasetsID sets the "datasets" edge to the Dataset entity by ID if the given value is not nil.
func (puo *ParameterUpdateOne) SetNillableDatasetsID(id *uuid.UUID) *ParameterUpdateOne {
	if id != nil {
		puo = puo.SetDatasetsID(*id)
	}
	return puo
}

// SetDatasets sets the "datasets" edge to the Dataset entity.
func (puo *ParameterUpdateOne) SetDatasets(d *Dataset) *ParameterUpdateOne {
	return puo.SetDatasetsID(d.ID)
}

// Mutation returns the ParameterMutation object of the builder.
func (puo *ParameterUpdateOne) Mutation() *ParameterMutation {
	return puo.mutation
}

// ClearDatasets clears the "datasets" edge to the Dataset entity.
func (puo *ParameterUpdateOne) ClearDatasets() *ParameterUpdateOne {
	puo.mutation.ClearDatasets()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ParameterUpdateOne) Select(field string, fields ...string) *ParameterUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Parameter entity.
func (puo *ParameterUpdateOne) Save(ctx context.Context) (*Parameter, error) {
	return withHooks[*Parameter, ParameterMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ParameterUpdateOne) SaveX(ctx context.Context) *Parameter {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ParameterUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ParameterUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ParameterUpdateOne) sqlSave(ctx context.Context) (_node *Parameter, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   parameter.Table,
			Columns: parameter.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: parameter.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Parameter.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, parameter.FieldID)
		for _, f := range fields {
			if !parameter.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != parameter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(parameter.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Unit(); ok {
		_spec.SetField(parameter.FieldUnit, field.TypeString, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(parameter.FieldCreatedAt, field.TypeTime, value)
	}
	if puo.mutation.DatasetsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DatasetsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Parameter{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{parameter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
