package avl_tree

func (t *Tree) PreOrderTraverse() []int {
	if t.root == nil {
		return []int{}
	} else {
		return t.preOrderTraverseFrom(t.root)
	}
}

func (t *Tree) preOrderTraverseFrom(node *Node) []int {
	result := make([]int, 0)
	result = append(result, node.value)
	if node.left != nil {
		result = append(result, t.preOrderTraverseFrom(node.left)...)
	}
	if node.right != nil {
		result = append(result, t.preOrderTraverseFrom(node.right)...)
	}
	return result
}
