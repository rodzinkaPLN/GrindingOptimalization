package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/db"
)

type Server struct {
	Port         int  `env:"PORT" envDefault:"8080"`
	RedirectMode bool `env:"REDIRECT_MODE" envDefault:"false"`
}

func (s *Server) Address() string {
	return fmt.Sprintf(":%d", s.Port)
}

type Config struct {
	Server Server
	DB     db.Config
}

func FromEnv() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("env parsing: %w", err)
	}
	return cfg, nil
}
