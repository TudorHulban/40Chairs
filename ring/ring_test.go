package ring

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateNodes(t *testing.T) {
	nodes := createNodes(1, 2)
	require.NotNil(t, nodes)
	require.Equal(t, nodes[0].ID, 1)
	require.Equal(t, nodes[1].ID, 2)
}

func TestRing(t *testing.T) {
	r := NewRing(1, 2)
	require.Equal(t, 2, len(r.Nodes))

	load := r.getAssignments()
	require.Equal(t, 2, len(load))

	for _, node := range load {
		assert.Greater(t, len(node.Load), 0, fmt.Sprintf("no load for node ID: %d", node.ID))
	}

	_, errAs := load.WriteTo(os.Stdout)
	require.NoError(t, errAs)
}
