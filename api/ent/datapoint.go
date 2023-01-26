// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/datapoint"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/parameter"
)

// Datapoint is the model entity for the Datapoint schema.
type Datapoint struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// DataTime holds the value of the "data_time" field.
	DataTime time.Time `json:"data_time,omitempty"`
	// ParameterID holds the value of the "parameter_id" field.
	ParameterID uuid.UUID `json:"parameter_id,omitempty"`
	// Val holds the value of the "val" field.
	Val float64 `json:"val,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DatapointQuery when eager-loading is set.
	Edges DatapointEdges `json:"edges"`
}

// DatapointEdges holds the relations/edges for other nodes in the graph.
type DatapointEdges struct {
	// Parameters holds the value of the parameters edge.
	Parameters *Parameter `json:"parameters,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParametersOrErr returns the Parameters value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DatapointEdges) ParametersOrErr() (*Parameter, error) {
	if e.loadedTypes[0] {
		if e.Parameters == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: parameter.Label}
		}
		return e.Parameters, nil
	}
	return nil, &NotLoadedError{edge: "parameters"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Datapoint) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case datapoint.FieldVal:
			values[i] = new(sql.NullFloat64)
		case datapoint.FieldDataTime, datapoint.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case datapoint.FieldID, datapoint.FieldParameterID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Datapoint", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Datapoint fields.
func (d *Datapoint) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case datapoint.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case datapoint.FieldDataTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field data_time", values[i])
			} else if value.Valid {
				d.DataTime = value.Time
			}
		case datapoint.FieldParameterID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field parameter_id", values[i])
			} else if value != nil {
				d.ParameterID = *value
			}
		case datapoint.FieldVal:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field val", values[i])
			} else if value.Valid {
				d.Val = value.Float64
			}
		case datapoint.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryParameters queries the "parameters" edge of the Datapoint entity.
func (d *Datapoint) QueryParameters() *ParameterQuery {
	return NewDatapointClient(d.config).QueryParameters(d)
}

// Update returns a builder for updating this Datapoint.
// Note that you need to call Datapoint.Unwrap() before calling this method if this Datapoint
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Datapoint) Update() *DatapointUpdateOne {
	return NewDatapointClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Datapoint entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Datapoint) Unwrap() *Datapoint {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Datapoint is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Datapoint) String() string {
	var builder strings.Builder
	builder.WriteString("Datapoint(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("data_time=")
	builder.WriteString(d.DataTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("parameter_id=")
	builder.WriteString(fmt.Sprintf("%v", d.ParameterID))
	builder.WriteString(", ")
	builder.WriteString("val=")
	builder.WriteString(fmt.Sprintf("%v", d.Val))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Datapoints is a parsable slice of Datapoint.
type Datapoints []*Datapoint

func (d Datapoints) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}