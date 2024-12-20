package cache

import (
	"errors"
)

var (
	ErrorUnloadCacheProvider = errors.New("unable to load a cache provider")
)

var cacheProviders = map[CacheProvider]func(config *CacheConfiguration) (ICache, error){
	REDIS_CACHE:  NewRedisProvider,
	MEMORY_CACHE: NewMemoryProvider,
}

func ConfigureCacheProvider(config *CacheConfiguration) (ICache, error) {

	p := CacheProvider(config.CacheProvider)

	callback, found := cacheProviders[p]

	if !found {
		return nil, ErrorUnloadCacheProvider
	}

	provider, err := callback(config)

	return provider, err
}
