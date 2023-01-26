package db

import (
	"database/sql"
	"embed"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/pressly/goose/v3"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/ent"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type Config struct {
	URL string `env:"DATABASE_URL,required" envDefault:"user=postgres host=localhost port=5432 dbname=postgres"`
}

type DB struct {
	*ent.Client
}

func NewWithMigrate(cfg Config) (*DB, error) {
	db, err := sql.Open("pgx", cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("open db failed: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping failed: %w", err)
	}
	db.SetMaxIdleConns(21)
	db.SetMaxOpenConns(37)

	goose.SetBaseFS(migrationsFS)

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, fmt.Errorf("database: can't set dialect %w", err)
	}

	if err := goose.Up(db, "migrations", goose.WithAllowMissing()); err != nil {
		return nil, fmt.Errorf("migrations up failed: %w", err)
	}

	client := ent.NewClient(
		ent.Driver(
			entsql.OpenDB(
				dialect.Postgres, db,
			),
		),
	)

	return &DB{
		Client: client,
	}, nil
}
