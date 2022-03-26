package ring

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRing(t *testing.T) {
	r := NewRing(1, 2)

	_, errAs := r.getAssignments().WriteTo(os.Stdout)
	require.NoError(t, errAs)
}
