package model

import (
	"time"

	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent"
)

type Datapoint struct {
	Ts  time.Time `json:"ts"`
	Val float64   `json:"v"`
}

type Parameter struct {
	Unit string      `json:"unit"`
	Name string      `json:"name"`
	Data []Datapoint `json:"data"`
}

func ParametersFromEntParameters(ins []*ent.Parameter) []Parameter {
	out := make([]Parameter, len(ins))

	for i, v := range ins {
		out[i] = Parameter{
			Name: v.Name,
			Unit: v.Unit,
			Data: make([]Datapoint, len(v.Edges.Datapoints)),
		}
		for j, vd := range v.Edges.Datapoints {
			out[i].Data[j] = Datapoint{
				Ts:  vd.DataTime,
				Val: vd.Val,
			}
		}
	}

	return out
}
