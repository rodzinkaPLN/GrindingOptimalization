// Code generated by ent, DO NOT EDIT.

package parameter

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the parameter type in the database.
	Label = "parameter"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDatasetID holds the string denoting the dataset_id field in the database.
	FieldDatasetID = "dataset_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeDatasets holds the string denoting the datasets edge name in mutations.
	EdgeDatasets = "datasets"
	// EdgeDatapoints holds the string denoting the datapoints edge name in mutations.
	EdgeDatapoints = "datapoints"
	// Table holds the table name of the parameter in the database.
	Table = "parameters"
	// DatasetsTable is the table that holds the datasets relation/edge.
	DatasetsTable = "parameters"
	// DatasetsInverseTable is the table name for the Dataset entity.
	// It exists in this package in order to avoid circular dependency with the "dataset" package.
	DatasetsInverseTable = "datasets"
	// DatasetsColumn is the table column denoting the datasets relation/edge.
	DatasetsColumn = "dataset_id"
	// DatapointsTable is the table that holds the datapoints relation/edge.
	DatapointsTable = "datapoints"
	// DatapointsInverseTable is the table name for the Datapoint entity.
	// It exists in this package in order to avoid circular dependency with the "datapoint" package.
	DatapointsInverseTable = "datapoints"
	// DatapointsColumn is the table column denoting the datapoints relation/edge.
	DatapointsColumn = "parameter_id"
)

// Columns holds all SQL columns for parameter fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDatasetID,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)