package ring

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateNodes(t *testing.T) {
	nodes := newNodes(1, 2)
	require.NotNil(t, nodes)
	require.Equal(t, nodes[0].ID, 1)
	require.Equal(t, nodes[1].ID, 2)
}

func TestRing(t *testing.T) {
	r := NewRing(1, 2)
	require.Greater(t, len(r.Nodes), 0)

	maxFactor := 5

	for i := 3; i < 40; i++ {
		r.RegisterNode(i)

		for f := 1; f <= maxFactor; f++ {
			r.SetFactor(f)

			require.GreaterOrEqual(t, len(r.Nodes), len(r.Assignments))
			require.Greater(t, len(r.Assignments), 0, "no assignments")

			require.NoError(t, r.verifyAssignments(), "self verification")
		}
	}

	_, errAs := r.Assignments.WriteTo(os.Stdout)
	require.NoError(t, errAs, "write to")
}
