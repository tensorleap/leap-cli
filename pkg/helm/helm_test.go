package helm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTensorleapChartValues(t *testing.T) {
	t.Run("CreateTensorleapChartValues with useGpu", func(t *testing.T) {

		useGpu := true
		expected := Record{
			"tensorleap-engine": Record{
				"gpu": useGpu,
			},
		}

		result := CreateTensorleapChartValues(useGpu)

		assert.Equal(t, expected, result)
	})
}
