package d2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTree_FindNearestToPoint(t *testing.T) {
	cases := map[string]struct {
		point Point
		want  Point
	}{
		"{5,5} want {5,4}":      {Point{5, 5}, Point{5, 4}},
		"{3, 4} want {3, 5}":    {Point{3, 4}, Point{3, 5}},
		"{3, 5} want {3, 5}":    {Point{3, 5}, Point{3, 5}},
		"{9, 9} want {3, 5}":    {Point{9, 9}, Point{9, 6}},
		"{7, 5} want {7, 4}":    {Point{7, 5}, Point{7, 4}},
		"{20, 20} want {9, 12}": {Point{20, 20}, Point{12, 6}},
	}

	points := []Point{{1, 2}, {9, 2}, {2, 5}, {3, 5}, {5, 4}, {7, 4}, {9, 6}, {12, 6}}
	tree := Tree{
		root: Build(points, 0),
	}

	tree.Dump()

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := tree.FindNearestToPoint(tc.point)
			require.Equal(t, tc.want, result)
		})
	}
}
