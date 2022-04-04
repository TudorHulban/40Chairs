package ring

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateNodes(t *testing.T) {
	nodes := newNodes(1, 2, 3)
	require.NotNil(t, nodes)
	require.Equal(t, nodes[0].ID, 1)
	require.Equal(t, nodes[1].ID, 2)
}

func TestDeRegisterNode(t *testing.T) {
	r := NewRing(1, 2, 3, 4)
	require.Equal(t, 4, len(r.Nodes))

	r.UnRegisterNode(2)
	require.Equal(t, []int{1, 3, 4}, r.Nodes.getIDs())
}

func TestAssignments(t *testing.T) {
	r := NewRing(1, 2)
	r.SetFactor(2)

	require.Equal(t, 20, len(r.Assignments.getRanges()))
}

func TestRing(t *testing.T) {
	r := NewRing(1)
	require.Greater(t, len(r.Nodes), 0)

	maxNodes := 3
	maxFactor := 2

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

func TestRedistribution(t *testing.T) {
	r := NewRing(1, 2, 3)
	require.Equal(t, 3, len(r.Nodes))

	r.SetFactor(2)

	assignments, errAs := os.Create("assignment_distribution_3_nodes")
	if errAs != nil {
		t.FailNow()
	}

	_, errWriteAssignments := r.Assignments.WriteTo(assignments)
	require.NoError(t, errWriteAssignments, "write to")

	r.UnRegisterNode(2)
	require.Equal(t, 2, len(r.Nodes))
	require.NoError(t, r.verifyAssignments(), "self verification")

	redistribution, errRed := os.Create("assignment_redistribution_2_nodes")
	if errRed != nil {
		t.FailNow()
	}

	_, errWriteRedistribution := r.Assignments.WriteTo(redistribution)
	require.NoError(t, errWriteRedistribution, "write to")
}

func TestUnregister40(t *testing.T) {

}
