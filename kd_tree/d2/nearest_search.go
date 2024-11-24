package d2

import (
	"math"
)

func distance(a Point, b Point) float64 {
	return math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2)
}

func (t *Tree) FindNearestToPoint(point Point) Point {
	best := &nd{}
	t.findNearestInSubtree(point, t.root, best)
	return best.node.point
}

type nd struct {
	node     *Node
	distance float64
}

func (t *Tree) findNearestInSubtree(point Point, node *Node, best *nd) {
	if node == nil {
		return
	}

	if node.dimension == DimensionNs {
		distance := distance(point, node.point)
		if best.node == nil || distance < best.distance {
			*best = nd{node: node, distance: distance}
		}
		return
	}

	var diff int

	if node.dimension == DimensionX {
		diff = point.x - node.value
	}

	if node.dimension == DimensionY {
		diff = point.y - node.value
	}

	if diff < 0 {
		t.findNearestInSubtree(point, node.left, best)
	} else {
		t.findNearestInSubtree(point, node.right, best)
	}

	var linearDistance float64
	if node.dimension == DimensionX {
		linearDistance = math.Abs(float64(point.x - node.value))
	} else {
		linearDistance = math.Abs(float64(point.y - node.value))
	}

	if best.node != nil && linearDistance < best.distance {
		// visit another branch
		if diff < 0 {
			t.findNearestInSubtree(point, node.right, best)
		} else {
			t.findNearestInSubtree(point, node.left, best)
		}
	}
}
