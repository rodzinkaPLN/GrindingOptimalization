package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/db"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/parameter"
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
	Parameters []string  `query:"parameters"`
	From       time.Time `query:"from"`
	To         time.Time `query:"to"`
}

type responseGetDatapoints struct {
	Parameters []model.Parameter `json:"parameters"`
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

	parameters, err := h.db.Parameter.Query().Where(
		parameter.NameIn(req.Parameters...),
	).WithDatapoints(
		func(q *ent.DatapointQuery) {
			q.Limit(100)
		},
	).
		All(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.DashboardFromEntParameters(parameters))
}
