package avl_tree

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Node struct {
	left   *Node
	right  *Node
	value  int
	height int
}

type Tree struct {
	root *Node
}
