package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	entDataset "github.com/rodzinkaPLN/GrindingOptimalization/api/ent/dataset"
)

func (h *CrudHandler) DataCSVUpload(c echo.Context) error {
	ctx := c.Request().Context()
	name := c.FormValue("dataset")

	dataset, err := h.db.Dataset.Query().Where(
		entDataset.NameEQ(name),
	).WithParameters().
		Only(ctx)
	if err != nil {
		return err
	}

	dataset.QueryParameters()

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// read csv file, column 0 is data, later on data points
	return c.NoContent(http.StatusNoContent)
}
