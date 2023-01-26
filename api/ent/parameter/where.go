// Code generated by ent, DO NOT EDIT.

package parameter

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Parameter {
	return predicate.Parameter(sql.FieldEQ(FieldName, v))
}

// DatasetID applies equality check predicate on the "dataset_id" field. It's identical to DatasetIDEQ.
func DatasetID(v uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldEQ(FieldDatasetID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Parameter {
	return predicate.Parameter(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Parameter {
	return predicate.Parameter(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Parameter {
	return predicate.Parameter(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Parameter {
	return predicate.Parameter(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Parameter {
	return predicate.Parameter(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Parameter {
	return predicate.Parameter(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Parameter {
	return predicate.Parameter(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Parameter {
	return predicate.Parameter(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Parameter {
	return predicate.Parameter(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Parameter {
	return predicate.Parameter(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Parameter {
	return predicate.Parameter(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Parameter {
	return predicate.Parameter(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Parameter {
	return predicate.Parameter(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Parameter {
	return predicate.Parameter(sql.FieldContainsFold(FieldName, v))
}

// DatasetIDEQ applies the EQ predicate on the "dataset_id" field.
func DatasetIDEQ(v uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldEQ(FieldDatasetID, v))
}

// DatasetIDNEQ applies the NEQ predicate on the "dataset_id" field.
func DatasetIDNEQ(v uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldNEQ(FieldDatasetID, v))
}

// DatasetIDIn applies the In predicate on the "dataset_id" field.
func DatasetIDIn(vs ...uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldIn(FieldDatasetID, vs...))
}

// DatasetIDNotIn applies the NotIn predicate on the "dataset_id" field.
func DatasetIDNotIn(vs ...uuid.UUID) predicate.Parameter {
	return predicate.Parameter(sql.FieldNotIn(FieldDatasetID, vs...))
}

// DatasetIDIsNil applies the IsNil predicate on the "dataset_id" field.
func DatasetIDIsNil() predicate.Parameter {
	return predicate.Parameter(sql.FieldIsNull(FieldDatasetID))
}

// DatasetIDNotNil applies the NotNil predicate on the "dataset_id" field.
func DatasetIDNotNil() predicate.Parameter {
	return predicate.Parameter(sql.FieldNotNull(FieldDatasetID))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Parameter {
	return predicate.Parameter(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Parameter {
	return predicate.Parameter(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Parameter {
	return predicate.Parameter(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Parameter {
	return predicate.Parameter(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Parameter {
	return predicate.Parameter(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Parameter {
	return predicate.Parameter(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Parameter {
	return predicate.Parameter(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Parameter {
	return predicate.Parameter(sql.FieldLTE(FieldCreatedAt, v))
}

// HasDatasets applies the HasEdge predicate on the "datasets" edge.
func HasDatasets() predicate.Parameter {
	return predicate.Parameter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DatasetsTable, DatasetsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDatasetsWith applies the HasEdge predicate on the "datasets" edge with a given conditions (other predicates).
func HasDatasetsWith(preds ...predicate.Dataset) predicate.Parameter {
	return predicate.Parameter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DatasetsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DatasetsTable, DatasetsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDatapoints applies the HasEdge predicate on the "datapoints" edge.
func HasDatapoints() predicate.Parameter {
	return predicate.Parameter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DatapointsTable, DatapointsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDatapointsWith applies the HasEdge predicate on the "datapoints" edge with a given conditions (other predicates).
func HasDatapointsWith(preds ...predicate.Datapoint) predicate.Parameter {
	return predicate.Parameter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DatapointsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DatapointsTable, DatapointsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Parameter) predicate.Parameter {
	return predicate.Parameter(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Parameter) predicate.Parameter {
	return predicate.Parameter(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Parameter) predicate.Parameter {
	return predicate.Parameter(func(s *sql.Selector) {
		p(s.Not())
	})
}