package config

import (
	"errors"

	"go.deanishe.net/env"
)

var (
	ErrorMarketDataProviderEnv      = errors.New("could not read UFOMS_CORE_MARKET_DATA_PROVIDER ENV")
	ErrorPositionServiceProviderEnv = errors.New("could not read UFOMS_CORE_POSITION_PROVIDER ENV")
	ErrorRiskServiceProviderEnv     = errors.New("could not read UFOMS_CORE_RISK_PROVIDER ENV")
)

type Configuration struct {
	CacheProvider             string `env:"UFOMS_CORE_CACHE_PROVIDER"`
	CachePrefix               string `env:"UFOMS_CORE_CACHE_PREFIX"`
	CacheHost                 string `env:"UFOMS_CORE_CACHE_HOST"`
	CacheDataBase             int    `env:"UFOMS_CORE_CACHE_DATA_BASE"`
	CachePassword             string `env:"UFOMS_CORE_CACHE_PASSWORD"`
	CacheUserName             string `env:"UFOMS_CORE_CACHE_USER_NAME"`
	CacheEnableTls            bool   `env:"UFOMS_CORE_CACHE_ENABLE_TLS"`
	CacheInsecureSkipVerify   bool   `env:"UFOMS_CORE_CACHE_INSECURE_SKIP_VERIFY"`
	CacheBufferSize           int    `env:"UFOMS_CORE_CACHE_BUFFER_SIZE"`
	NewOrderChannel           string `env:"UFOMS_CORE_NEW_ORDER_CHANNEL"`
	MarketDataServiceProvider string `env:"UFOMS_CORE_MARKET_DATA_PROVIDER"`
	PositionServiceProvider   string `env:"UFOMS_CORE_POSITION_PROVIDER"`
	RiskServiceProvider       string `env:"UFOMS_CORE_RISK_PROVIDER"`
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
	if c.MarketDataServiceProvider == "" {
		errs = append(errs, ErrorMarketDataProviderEnv)
	}

	if c.PositionServiceProvider == "" {
		errs = append(errs, ErrorPositionServiceProviderEnv)
	}

	if c.RiskServiceProvider == "" {
		errs = append(errs, ErrorRiskServiceProviderEnv)
	}

	return errs
}
