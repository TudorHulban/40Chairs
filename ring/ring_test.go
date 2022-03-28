package ring

import (
	"fmt"
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

func TestAssignments(t *testing.T) {
	r := NewRing(1, 2)
	r.SetFactor(2)

	require.Equal(t, 20, len(r.Assignments.getRanges()))
}

func TestRing(t *testing.T) {
	r := NewRing(1)
	require.Greater(t, len(r.Nodes), 0)

	maxNodes := 4
	maxFactor := 3

	output, errCr := os.Create("assignment_distribution")
	if errCr != nil {
		t.FailNow()
	}

	writeTo := output
	writeTo.WriteString(fmt.Sprintf("Load: %v.\n\n", r.Load.getUniqueRanges()))

	for i := 2; i <= maxNodes; i++ {
		r.RegisterNode(i)

		for f := 1; f <= maxFactor; f++ {
			if f > len(r.Nodes) {
				continue
			}

			r.SetFactor(f)

			require.GreaterOrEqual(t, len(r.Nodes), len(r.Assignments))
			require.Greater(t, len(r.Assignments), 0, "no assignments")

			require.NoError(t, r.verifyAssignments(), "self verification")

			writeTo.WriteString(fmt.Sprintf("Nodes: %d. Factor: %d\n", len(r.Nodes), f))
			// writeTo.WriteString(fmt.Sprintf("Assignments: %v.\n", r.Assignments.getRanges()))

			_, errAs := r.Assignments.WriteTo(writeTo)
			require.NoError(t, errAs, "write to")
		}
	}
}
