package k3d

import (
	"testing"

	k3d "github.com/k3d-io/k3d/v5/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestIsGpuCluster(t *testing.T) {
	testCases := []struct {
		name     string
		cluster  *Cluster
		expected bool
	}{
		{
			name: "Non-GPU cluster",
			cluster: &Cluster{
				Nodes: []*k3d.Node{
					{Image: K3sImage},
				},
			},
			expected: false,
		},
		{
			name: "GPU cluster",
			cluster: &Cluster{
				Nodes: []*k3d.Node{
					{Image: K3sGpuImage},
				},
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsGpuCluster(tc.cluster)
			assert.Equal(t, tc.expected, result)
		})
	}
}
