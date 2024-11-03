package database

import (
	"errors"

	"go.deanishe.net/env"
)

var (
	ErrorDBProviderEnv = errors.New("could not read UFOMS_DB_PROVIDER ENV")
)

type Configuration struct {
	DBProvider string `env:"UFOMS_DB_PROVIDER"`
}

func LoadEnvironments() (*Configuration, []error) {

	cfg := &Configuration{}
	err := env.Bind(cfg)

	if err != nil {
		return nil, []error{err}
	}

	errs := cfg.validateEnvironmentVariables()

	if len(errs) > 0 {
		return nil, errs
	}

	return cfg, nil
}

func (c *Configuration) validateEnvironmentVariables() []error {

	errs := make([]error, 0)

	if c.DBProvider == "" {
		errs = append(errs, ErrorDBProviderEnv)
	}

	return errs
}
