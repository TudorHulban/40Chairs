package ring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRanges(t *testing.T) {
	n1 := Node{
		ID:   1,
		Load: []Range{"1", "2"},
	}

	n2 := Node{
		ID:   2,
		Load: []Range{"3"},
	}

	nodes := Nodes{&n2, &n1}
	require.Equal(t, 2, len(nodes.getIDs()))
	require.Equal(t, []Range{"1", "2", "3"}, nodes.getRanges())
}
