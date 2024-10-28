package cache_test

import (
	"os"
	"testing"

	"github.com/jefferson1208/ufoms/pkg/cache"
	"github.com/stretchr/testify/assert"
)

func TestLoadEnvironments(t *testing.T) {

	t.Run("should fail to load the settings for cache", func(t *testing.T) {

		cfg, errs := cache.LoadEnvironments()

		assert.Nil(t, cfg)
		assert.NotNil(t, errs)

		assert.Equal(t, cache.ErrorCacheProviderEnv, errs[0])
		assert.Equal(t, cache.ErrorCacheHostEnv, errs[1])
		assert.Equal(t, cache.ErrorCacheDataBaseEnv, errs[2])
		assert.Equal(t, cache.ErrorCacheBufferSizeEnv, errs[3])

	})

	t.Run("should load the settings for cache based on the envs", func(t *testing.T) {

		os.Setenv("UFOMS_CACHE_PROVIDER", "REDIS")
		os.Setenv("UFOMS_CACHE_HOST", "redis:6379")
		os.Setenv("UFOMS_CACHE_DATA_BASE", "1")
		os.Setenv("UFOMS_CACHE_BUFFER_SIZE", "1000")

		cfg, errs := cache.LoadEnvironments()

		assert.Nil(t, errs)
		assert.NotNil(t, cfg)
	})

}
