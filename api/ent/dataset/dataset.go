// Code generated by ent, DO NOT EDIT.

package dataset

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the dataset type in the database.
	Label = "dataset"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeParameters holds the string denoting the parameters edge name in mutations.
	EdgeParameters = "parameters"
	// Table holds the table name of the dataset in the database.
	Table = "datasets"
	// ParametersTable is the table that holds the parameters relation/edge.
	ParametersTable = "parameters"
	// ParametersInverseTable is the table name for the Parameter entity.
	// It exists in this package in order to avoid circular dependency with the "parameter" package.
	ParametersInverseTable = "parameters"
	// ParametersColumn is the table column denoting the parameters relation/edge.
	ParametersColumn = "dataset_id"
)

// Columns holds all SQL columns for dataset fields.
var Columns = []string{
	FieldID,
	FieldName,
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