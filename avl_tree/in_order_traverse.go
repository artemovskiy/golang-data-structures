package avl_tree

func (t *Tree) SymmetricTraverse() []int {
	if t.root == nil {
		return []int{}
	} else {
		return t.symmetricTraverseFrom(t.root)
	}
}

func (t *Tree) symmetricTraverseFrom(node *Node) []int {

	result := make([]int, 0)
	if node.left != nil {
		result = append(result, t.symmetricTraverseFrom(node.left)...)
	}
	result = append(result, node.value)
	if node.right != nil {
		result = append(result, t.symmetricTraverseFrom(node.right)...)
	}
	return result
}
