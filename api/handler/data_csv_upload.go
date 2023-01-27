package handler

import (
	"context"
	"encoding/csv"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent"
	entDataset "github.com/rodzinkaPLN/GrindingOptimalization/api/ent/dataset"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent/schema"
)

const batchingNo = 4096

func (h *CrudHandler) DataCSVUpload(c echo.Context) error {
	ctx := c.Request().Context()
	name := c.FormValue("dataset")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(src)

	header, err := csvReader.Read()
	if err != nil {
		return err
	}

	tx, err := h.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	columns := header[2:]
	dataset, err := h.db.
		Dataset.
		Query().Where(
		entDataset.NameEQ(name),
	).WithParameters().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			dataset, err = tx.Dataset.
				Create().
				SetName(name).
				Save(ctx)
			if err != nil {
				return err
			}
		} else {
			return err
		}
		paramsCreate := make([]*ent.ParameterCreate, len(columns))
		for i, v := range columns {
			paramsCreate[i] = tx.
				Parameter.
				Create().
				SetDatasetID(dataset.ID).
				SetName(v).SetUnit("g")
		}
		params, err := tx.
			Parameter.
			CreateBulk(paramsCreate...).
			Save(ctx)

		if err != nil {
			return err
		}
		dataset.Edges.Parameters = params
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	go func() {
		ctx = context.Background()
		err := func() error {
			toInsert := make([]*ent.DatapointCreate, 0, batchingNo)
			for row, err := csvReader.Read(); err == nil; row, err = csvReader.Read() {
				date, err := time.Parse("2006-01-02 15:04:05", row[1])
				if err != nil {
					return err
				}

				d := make(schema.DataT)
				for i, v := range row[2:] {
					if v == "" {
						// skip empty values
						continue
					}
					fVal, err := strconv.ParseFloat(v, 64)
					if err != nil {
						return err
					}
					d[columns[i]] = fVal
				}
				toInsert = append(toInsert, h.db.
					Datapoint.
					Create().
					SetDatasetID(dataset.ID).
					SetVals(d).
					SetDataTime(date),
				)
				if len(toInsert) >= batchingNo {
					log.Println("insert batch")
					err = h.db.Datapoint.
						CreateBulk(toInsert...).
						Exec(ctx)
					if err != nil {
						return err
					}
					toInsert = make([]*ent.DatapointCreate, 0, batchingNo)
				}
			}
			err = h.db.Datapoint.
				CreateBulk(toInsert...).
				Exec(ctx)
			if err != nil {
				return err
			}
			return nil
		}()
		defer src.Close()
		if err != nil {
			log.Println("upload failed ", err)
		}
	}()

	// read csv file, column 0 is data, later on data points
	return c.NoContent(http.StatusNoContent)
}
