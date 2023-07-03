package helm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTensorleapChartValues(t *testing.T) {
	t.Run("CreateTensorleapChartValues", func(t *testing.T) {
		useGpu := true
		dataDir := "some/dir/path"
		expected := Record{
			"tensorleap-engine": Record{
				"gpu":                useGpu,
				"localDataDirectory": dataDir,
			},
		}

		result := CreateTensorleapChartValues(useGpu, dataDir)

		assert.Equal(t, expected, result)
	})
}
