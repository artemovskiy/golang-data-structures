package d2

import (
	"cmp"
	"slices"
)

func Build(points []Point, depth int) *Node {
	if len(points) == 0 {
		return nil
	}
	if len(points) == 1 {
		return &Node{
			point: points[0],
		}
	}
	sortedPoints := make([]Point, len(points))
	copy(sortedPoints, points)
	slices.SortFunc(sortedPoints, func(a, b Point) int {
		if depth%2 == 0 {
			return cmp.Compare(a.x, b.x)
		} else {
			return cmp.Compare(a.y, b.y)
		}
	})
	medianIndex := len(sortedPoints) / 2
	medianPoint := sortedPoints[medianIndex]

	dim := DimensionX
	value := medianPoint.x
	if depth%2 == 1 {
		dim = DimensionY
		value = medianPoint.y
	}

	node := Node{
		dimension: dim,
		value:     value,
	}

	node.left = Build(sortedPoints[:medianIndex], depth+1)
	node.right = Build(sortedPoints[medianIndex:], depth+1)

	return &node
}
