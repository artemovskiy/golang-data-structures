package binary_trees

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIterator(t *testing.T) {
	var result []Item
	for v := range iterateItems {
		result = append(result, v)
	}

	expect := []Item{1, 2, 3}

	require.Equal(t, expect, result)
}
