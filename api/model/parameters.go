package model

import (
	"time"

	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent"
)

type Parameter struct {
	Unit string `json:"unit"`
	Name string `json:"name"`
	Key  string `json:"key"`
}
type Dashboard struct {
	Parameters []Parameter              `json:"parameters"`
	Data       []map[string]interface{} `json:"data"`
}
type Prediction struct {
	Date time.Time      `json:"date"`
	Data map[string]any `json:"data"`
}

type Userinput struct {
	Min          float64 `json:"min"`
	Max          float64 `json:"max"`
	DefaultValue float64 `json:"default_value"`
	Step         float64 `json:"step"`
	Name         string  `json:"name"`
}
type Userinputs struct {
	Data []Userinput `json:"data"`
}

func DashboardFromEntDataset(in *ent.Dataset) Dashboard {
	out := Dashboard{
		Parameters: make([]Parameter, len(in.Edges.Parameters)),
		Data:       make([]map[string]interface{}, len(in.Edges.Datapoints)),
	}

	for i, v := range in.Edges.Parameters {
		out.Parameters[i] = Parameter{
			Unit: v.Unit,
			Name: v.Name,
			Key:  v.Name,
		}
	}
	for i, v := range in.Edges.Datapoints {
		out.Data[i] = make(map[string]interface{})
		out.Data[i]["Data"] = v.DataTime
		for key, val := range v.Vals {
			out.Data[i][key] = val
		}
	}
	return out
}

func PredictionFromEntPrediction(in *ent.Prediction) Prediction {
	var out Prediction
	out.Data = make(map[string]any)
	for k, v := range in.Vals {
		out.Data[k] = v
	}
	return out
}

func UserinputFromEntUserinput(ins []*ent.Userinput) Userinputs {
	var out Userinputs
	out.Data = make([]Userinput, len(ins))
	for k, v := range ins {
		out.Data[k] = Userinput{
			Min:          v.Min,
			Max:          v.Max,
			DefaultValue: v.Defaultval,
			Step:         v.Step,
			Name:         v.Name,
		}
	}
	return out
}
