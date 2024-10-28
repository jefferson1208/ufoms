package cache

import (
	"errors"

	"go.deanishe.net/env"
)

var (
	ErrorCacheProviderEnv = errors.New("could not read UFOMS_CACHE_PROVIDER ENV")
)

type Configuration struct {
	CacheProvider           string `env:"UFOMS_CACHE_PROVIDER"`
	CachePrefix             string `env:"UFOMS_CACHE_PREFIX"`
	CacheHost               string `env:"UFOMS_CACHE_HOST"`
	CacheDataBase           int    `env:"UFOMS_CACHE_DATA_BASE"`
	CachePassword           string `env:"UFOMS_CACHE_PASSWORD"`
	CacheUserName           string `env:"UFOMS_CACHE_USER_NAME"`
	CacheEnableTls          bool   `env:"UFOMS_CACHE_ENABLE_TLS"`
	CacheInsecureSkipVerify bool   `env:"UFOMS_CACHE_INSECURE_SKIP_VERIFY"`
	CacheBufferSize         int    `env:"UFOMS_CACHE_BUFFER_SIZE"`
	Channels                string `env:"UFOMS_NEW_ORDER_CHANNEL"`
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
	if c.CacheProvider == "" {
		errs = append(errs, ErrorCacheProviderEnv)
	}

	return errs
}
