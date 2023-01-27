package model

import (
	"sort"
	"strconv"
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
	Data       []map[string]interface{} `json:"data"	`
}

func DashboardFromEntParameters(ins []*ent.Parameter) Dashboard {
	out := Dashboard{}

	d := make(map[time.Time]map[string]float64)

	for i, v := range ins {
		key := strconv.Itoa(i)
		out.Parameters = append(out.Parameters, Parameter{
			Unit: v.Unit,
			Name: v.Name,
			Key:  key,
		})

		for _, vv := range v.Edges.Datapoints {
			if _, ok := d[vv.DataTime]; !ok {
				d[vv.DataTime] = make(map[string]float64)
			}
			d[vv.DataTime][key] = vv.Val
		}
	}

	dd := make([]time.Time, 0, len(d))
	for k := range d {
		dd = append(dd, k)
	}
	sort.Slice(dd, func(i, j int) bool {
		return dd[i].Before(dd[j])
	})

	out.Data = make([]map[string]interface{}, len(dd))
	for i, t := range dd {
		dataCol := make(map[string]interface{})
		dataCol["ts"] = t

		for ii, v := range d[t] {
			dataCol[ii] = v
		}

		out.Data[i] = dataCol
	}

	return out
}
