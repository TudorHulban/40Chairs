package ring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFactoredRanges(t *testing.T) {
	p20 := NewPartitions(2)

	require.Equal(t, 20, len(p20.getFactoredRanges()))
}
