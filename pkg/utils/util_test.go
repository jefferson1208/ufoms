package utils_test

import (
	"testing"

	"github.com/jefferson1208/ufoms/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestConvertFromAscii(t *testing.T) {

	t.Run("must convert to zero", func(t *testing.T) {

		converted := utils.ConvertFromAscii(48)

		assert.Equal(t, "0", converted)
	})

	t.Run("must convert to one", func(t *testing.T) {

		converted := utils.ConvertFromAscii(49)

		assert.Equal(t, "1", converted)
	})

	t.Run("must convert to two", func(t *testing.T) {

		converted := utils.ConvertFromAscii(50)

		assert.Equal(t, "2", converted)
	})

	t.Run("must convert to three", func(t *testing.T) {

		converted := utils.ConvertFromAscii(51)

		assert.Equal(t, "3", converted)
	})

	t.Run("must convert to four", func(t *testing.T) {

		converted := utils.ConvertFromAscii(52)

		assert.Equal(t, "4", converted)
	})

	t.Run("must convert to five", func(t *testing.T) {

		converted := utils.ConvertFromAscii(53)

		assert.Equal(t, "5", converted)
	})

	t.Run("must convert to six", func(t *testing.T) {

		converted := utils.ConvertFromAscii(54)

		assert.Equal(t, "6", converted)
	})

	t.Run("must convert to seven", func(t *testing.T) {

		converted := utils.ConvertFromAscii(55)

		assert.Equal(t, "7", converted)
	})

	t.Run("must convert to eight", func(t *testing.T) {

		converted := utils.ConvertFromAscii(56)

		assert.Equal(t, "8", converted)
	})

	t.Run("must convert to nine", func(t *testing.T) {

		converted := utils.ConvertFromAscii(57)

		assert.Equal(t, "9", converted)
	})

}
