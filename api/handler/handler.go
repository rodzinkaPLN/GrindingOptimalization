package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/db"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/datapoint"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/dataset"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/prediction"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/userinput"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/model"
)

type CrudHandler struct {
	db *db.DB
}

func NewCrudHandler(dbc *db.DB) *CrudHandler {
	return &CrudHandler{
		db: dbc,
	}
}

type requestGetDatapoints struct {
	Dataset string    `query:"dataset"`
	From    time.Time `query:"from"`
	To      time.Time `query:"to"`
}

func (h *CrudHandler) GetDataPoints(c echo.Context) error {
	ctx := c.Request().Context()
	var req requestGetDatapoints
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	dataset, err := h.db.Dataset.Query().Where(
		dataset.NameEQ(req.Dataset),
	).WithDatapoints(
		func(q *ent.DatapointQuery) {
			q.Where(
				datapoint.And(
					datapoint.DataTimeGTE(req.From),
					datapoint.DataTimeLTE(req.To),
				),
			)
			q.Limit(1500)
			q.Order(ent.Asc(datapoint.FieldDataTime))
		},
	).WithParameters().
		First(ctx)
	fmt.Println(dataset.Edges.Parameters)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.DashboardFromEntDataset(dataset))
}

func (h *CrudHandler) GetPredictions(c echo.Context) error {
	ctx := c.Request().Context()

	prediction, err := h.db.
		Prediction.
		Query().
		Order(ent.Desc(prediction.FieldDataTime)).First(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.PredictionFromEntPrediction(prediction))
}

func (h *CrudHandler) GetUserinputs(c echo.Context) error {
	ctx := c.Request().Context()

	userinputs, err := h.db.
		Userinput.
		Query().
		Order(ent.Desc(userinput.FieldName)).All(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.UserinputFromEntUserinput(userinputs))
}

type requestGetUsage struct {
	Water float64 `query:"water"`
	Cu    float64 `query:"cu"`
	Speed float64 `query:"speed"`
}

type UsageData struct {
	Energy    float64 `json:"energy"`
	Granulate float64 `json:"granulate"`
}
type responseGetUsage struct {
	Data UsageData `json:"data"`
}

func (h *CrudHandler) GetUsage(c echo.Context) error {
	var req requestGetUsage
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}
	// TODO: replace with datamodel script usage

	time.Sleep(2 * time.Second)
	return c.JSON(http.StatusOK, responseGetUsage{
		UsageData{
			Energy: 1000 + req.Cu*0.02 +
				req.Speed*1.05 +
				req.Water*0.3,
			Granulate: 200 + req.Cu*0.72 +
				req.Speed*1.75 +
				req.Water*0.02,
		},
	})
}
