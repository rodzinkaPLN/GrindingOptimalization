package handler

import (
	"log"

	"github.com/labstack/echo/v4"
)

type responseListUserFiles struct {
}

func (h *CrudHandler) DataCSVUpload(c echo.Context) error {
	ctx := c.Request().Context()
	uid, err := getUID(c)
	if err != nil {
		return err
	}
	log.Print(uid, ctx)
	return nil
}
