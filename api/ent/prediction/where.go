// Code generated by ent, DO NOT EDIT.

package prediction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Prediction {
	return predicate.Prediction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Prediction {
	return predicate.Prediction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Prediction {
	return predicate.Prediction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Prediction {
	return predicate.Prediction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Prediction {
	return predicate.Prediction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Prediction {
	return predicate.Prediction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Prediction {
	return predicate.Prediction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Prediction {
	return predicate.Prediction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Prediction {
	return predicate.Prediction(sql.FieldLTE(FieldID, id))
}

// DataTime applies equality check predicate on the "data_time" field. It's identical to DataTimeEQ.
func DataTime(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldEQ(FieldDataTime, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldEQ(FieldCreatedAt, v))
}

// DataTimeEQ applies the EQ predicate on the "data_time" field.
func DataTimeEQ(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldEQ(FieldDataTime, v))
}

// DataTimeNEQ applies the NEQ predicate on the "data_time" field.
func DataTimeNEQ(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldNEQ(FieldDataTime, v))
}

// DataTimeIn applies the In predicate on the "data_time" field.
func DataTimeIn(vs ...time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldIn(FieldDataTime, vs...))
}

// DataTimeNotIn applies the NotIn predicate on the "data_time" field.
func DataTimeNotIn(vs ...time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldNotIn(FieldDataTime, vs...))
}

// DataTimeGT applies the GT predicate on the "data_time" field.
func DataTimeGT(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldGT(FieldDataTime, v))
}

// DataTimeGTE applies the GTE predicate on the "data_time" field.
func DataTimeGTE(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldGTE(FieldDataTime, v))
}

// DataTimeLT applies the LT predicate on the "data_time" field.
func DataTimeLT(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldLT(FieldDataTime, v))
}

// DataTimeLTE applies the LTE predicate on the "data_time" field.
func DataTimeLTE(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldLTE(FieldDataTime, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Prediction {
	return predicate.Prediction(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Prediction) predicate.Prediction {
	return predicate.Prediction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Prediction) predicate.Prediction {
	return predicate.Prediction(func(s *sql.Selector) {
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
func Not(p predicate.Prediction) predicate.Prediction {
	return predicate.Prediction(func(s *sql.Selector) {
		p(s.Not())
	})
}
