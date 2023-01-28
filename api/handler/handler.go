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
