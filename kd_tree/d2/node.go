package d2

type Point struct {
	x int
	y int
}

type Dimension string

const (
	DimensionNs Dimension = ""
	DimensionX  Dimension = "x"
	DimensionY  Dimension = "y"
)

type Node struct {
	point     Point
	dimension Dimension
	value     int
	left      *Node
	right     *Node
}

type Tree struct {
	root *Node
}
