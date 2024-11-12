package cache

import (
	"errors"
)

var (
	ErrorCacheProviderEnv   = errors.New("could not read UFOMS_CACHE_PROVIDER ENV")
	ErrorCacheHostEnv       = errors.New("could not read UFOMS_CACHE_HOST ENV")
	ErrorCacheDataBaseEnv   = errors.New("could not read UFOMS_CACHE_DATA_BASE ENV")
	ErrorCacheBufferSizeEnv = errors.New("could not read UFOMS_CACHE_BUFFER_SIZE ENV")
)

type CacheConfiguration struct {
	CacheProvider           string
	CachePrefix             string
	CacheHost               string
	CacheDataBase           int
	CachePassword           string
	CacheUserName           string
	CacheEnableTls          bool
	CacheInsecureSkipVerify bool
	CacheBufferSize         int
	Channels                string
}

func Load(c *CacheConfiguration) (*CacheConfiguration, []error) {

	errs := c.validateEnvironmentVariables()

	if len(errs) > 0 {
		return nil, errs
	}

	return c, nil
}

func (c *CacheConfiguration) validateEnvironmentVariables() []error {

	errs := make([]error, 0)

	if c.CacheProvider == "" {
		errs = append(errs, ErrorCacheProviderEnv)
	}

	if c.CacheHost == "" {
		errs = append(errs, ErrorCacheHostEnv)
	}

	if c.CacheDataBase < 1 {
		errs = append(errs, ErrorCacheDataBaseEnv)
	}

	if c.CacheBufferSize < 1 {
		errs = append(errs, ErrorCacheBufferSizeEnv)
	}

	return errs
}
