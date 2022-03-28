package ring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteAtIndex(t *testing.T) {
	s := []int{1, 2, 3, 4}

	deleteAtIndex[int](&s, 1)
	require.Equal(t, []int{1, 3, 4}, s)
}
