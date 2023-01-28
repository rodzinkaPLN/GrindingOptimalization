// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/datapoint"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/dataset"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/parameter"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/prediction"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	datapointFields := schema.Datapoint{}.Fields()
	_ = datapointFields
	// datapointDescCreatedAt is the schema descriptor for created_at field.
	datapointDescCreatedAt := datapointFields[4].Descriptor()
	// datapoint.DefaultCreatedAt holds the default value on creation for the created_at field.
	datapoint.DefaultCreatedAt = datapointDescCreatedAt.Default.(func() time.Time)
	// datapointDescID is the schema descriptor for id field.
	datapointDescID := datapointFields[0].Descriptor()
	// datapoint.DefaultID holds the default value on creation for the id field.
	datapoint.DefaultID = datapointDescID.Default.(func() uuid.UUID)
	datasetFields := schema.Dataset{}.Fields()
	_ = datasetFields
	// datasetDescCreatedAt is the schema descriptor for created_at field.
	datasetDescCreatedAt := datasetFields[2].Descriptor()
	// dataset.DefaultCreatedAt holds the default value on creation for the created_at field.
	dataset.DefaultCreatedAt = datasetDescCreatedAt.Default.(func() time.Time)
	// datasetDescID is the schema descriptor for id field.
	datasetDescID := datasetFields[0].Descriptor()
	// dataset.DefaultID holds the default value on creation for the id field.
	dataset.DefaultID = datasetDescID.Default.(func() uuid.UUID)
	parameterFields := schema.Parameter{}.Fields()
	_ = parameterFields
	// parameterDescCreatedAt is the schema descriptor for created_at field.
	parameterDescCreatedAt := parameterFields[4].Descriptor()
	// parameter.DefaultCreatedAt holds the default value on creation for the created_at field.
	parameter.DefaultCreatedAt = parameterDescCreatedAt.Default.(func() time.Time)
	// parameterDescID is the schema descriptor for id field.
	parameterDescID := parameterFields[0].Descriptor()
	// parameter.DefaultID holds the default value on creation for the id field.
	parameter.DefaultID = parameterDescID.Default.(func() uuid.UUID)
	predictionFields := schema.Prediction{}.Fields()
	_ = predictionFields
	// predictionDescCreatedAt is the schema descriptor for created_at field.
	predictionDescCreatedAt := predictionFields[3].Descriptor()
	// prediction.DefaultCreatedAt holds the default value on creation for the created_at field.
	prediction.DefaultCreatedAt = predictionDescCreatedAt.Default.(func() time.Time)
	// predictionDescID is the schema descriptor for id field.
	predictionDescID := predictionFields[0].Descriptor()
	// prediction.DefaultID holds the default value on creation for the id field.
	prediction.DefaultID = predictionDescID.Default.(func() uuid.UUID)
}
