package d2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTree_Build(t *testing.T) {
	cases := map[string]struct {
		points []Point
	}{
		"first": {
			points: []Point{{1, 2}, {9, 2}, {2, 5}, {3, 5}, {5, 4}, {7, 4}, {9, 6}, {12, 6}},
		},
	}

	for k, c := range cases {
		t.Run(k, func(t *testing.T) {
			tree := Tree{
				root: Build(c.points, 0),
			}

			tree.Dump()
			require.Equal(t, 1, 1)
		})
	}
}
