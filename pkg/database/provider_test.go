package database_test

import (
	"testing"

	"github.com/jefferson1208/ufoms/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestConfigureDBProvider(t *testing.T) {

	t.Run("should return the MYSQL provider", func(t *testing.T) {

		cfg := &database.Configuration{DBProvider: "MYSQL"}
		provider, err := database.ConfigureDBProvider(cfg)

		assert.Nil(t, err)
		assert.NotNil(t, provider)
		assert.Equal(t, string(database.MYSQL_DB), provider.GetProvider())

	})

	t.Run("should return the MEMORY provider", func(t *testing.T) {

		cfg := &database.Configuration{DBProvider: "MEMORY"}
		provider, err := database.ConfigureDBProvider(cfg)

		assert.Nil(t, err)
		assert.NotNil(t, provider)
		assert.Equal(t, string(database.MEMORY_DB), provider.GetProvider())

	})

	t.Run("should not create a cache provider", func(t *testing.T) {

		cfg := &database.Configuration{DBProvider: "XPTO"}
		provider, err := database.ConfigureDBProvider(cfg)

		assert.Nil(t, provider)
		assert.NotNil(t, err)
		assert.Equal(t, database.ErrorUnloadDBProvider, err)

	})

}
