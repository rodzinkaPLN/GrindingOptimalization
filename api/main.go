package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/config"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/db"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/handler"
)

type V struct {
	Validator *validator.Validate
}

func (cv *V) Validate(i any) error {
	return cv.Validator.Struct(i)
}

func main() {

	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatal(err)
	}
	if cfg.Server.RedirectMode {
		e := echo.New()
		e.Any("*", func(c echo.Context) error {
			return c.Redirect(http.StatusPermanentRedirect, "https://FIXME")
		})
		log.Fatal(e.Start(cfg.Server.Address()))
	}

	db, err := db.NewWithMigrate(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	crudHandler := handler.NewCrudHandler(db)

	log.Print(crudHandler)

	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.CORS(),
		middleware.Logger(),
	)

	e.Validator = &V{validator.New()}
	api := e.Group("/api/v1")
	api.GET("/healthcheck", func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, []byte(`{"status":"ok"}`))
	})

	// static for react
	staticPath := "/app/public"
	err = filepath.Walk(staticPath,
		func(path string, _ os.FileInfo, _ error) error {
			routePath := path[len(staticPath):]
			e.GET(routePath, func(c echo.Context) error {
				return c.File(path)
			})
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	e.Any("*", func(c echo.Context) error {
		return c.File("/app/public/index.html")
	})

	data := api.Group("/data")
	data.POST(":csv-upload", crudHandler.DataCSVUpload)
	data.GET("", crudHandler.GetDataPoints)

	log.Fatal(e.Start(cfg.Server.Address()))
}
