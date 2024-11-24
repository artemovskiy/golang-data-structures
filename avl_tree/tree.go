package avl_tree

type Node struct {
	left   *Node
	right  *Node
	value  int
	height int
}

type Tree struct {
	root *Node
}
