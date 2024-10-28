package metrics_test

import (
	"testing"

	"github.com/jefferson1208/ufoms/pkg/metrics"
	"github.com/stretchr/testify/assert"
)

func TestNewMetricHandler(t *testing.T) {

	t.Run("should fail to load the settings", func(t *testing.T) {

		metricHandler, err := metrics.NewMetricHandler("CONSOLE", 100)

		assert.Nil(t, err)
		assert.NotNil(t, metricHandler)

		assert.Equal(t, string(metrics.CONSOLE_OUTPUT), metricHandler.GetOutputName())
	})
}
