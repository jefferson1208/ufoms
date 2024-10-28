package cache_test

import (
	"testing"

	"github.com/jefferson1208/ufoms/pkg/cache"
	"github.com/stretchr/testify/assert"
)

func TestConfigureCacheProvider(t *testing.T) {

	t.Run("should return the REDIS provider", func(t *testing.T) {

		cfg := &cache.Configuration{CacheProvider: "REDIS"}
		provider, err := cache.ConfigureCacheProvider(cfg)

		assert.Nil(t, err)
		assert.NotNil(t, provider)
		assert.Equal(t, string(cache.REDIS_CACHE), provider.GetProvider())

	})

	t.Run("should return the MEMORY provider", func(t *testing.T) {

		cfg := &cache.Configuration{CacheProvider: "MEMORY"}
		provider, err := cache.ConfigureCacheProvider(cfg)

		assert.Nil(t, err)
		assert.NotNil(t, provider)
		assert.Equal(t, string(cache.MEMORY_CACHE), provider.GetProvider())

	})

	t.Run("should not create a cache provider", func(t *testing.T) {

		cfg := &cache.Configuration{CacheProvider: "xpto"}
		provider, err := cache.ConfigureCacheProvider(cfg)

		assert.Nil(t, provider)
		assert.NotNil(t, err)
		assert.Equal(t, cache.ErrorUnloadCacheProvider, err)

	})

}
