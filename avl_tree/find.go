package avl_tree

func (t *Tree) Find(v int) bool {
	node := t.root
	for node != nil {
		if node.value == v {
			return true
		}
		if node.value > v {
			node = node.left
		}
		if node.value < v {
			node = node.right
		}
	}
	return false
}
