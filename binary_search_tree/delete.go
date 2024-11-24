package binary_search_tree

func (t *Tree) Delete(v int) {
	t.root = t.deleteFromSubtree(v, t.root)
}

func (t *Tree) deleteFromSubtree(v int, parent *Node) *Node {
	if parent == nil {
		return nil
	}
	if v < parent.value {
		parent.left = t.deleteFromSubtree(v, parent.left)
	} else if v > parent.value {
		parent.right = t.deleteFromSubtree(v, parent.right)
	} else {
		if parent.left == nil && parent.right == nil {
			return nil
		}
		if parent.left == nil {
			return parent.right
		}
		if parent.right == nil {
			return parent.left
		}

		// find the least node in the right subtree of the current node
		node := parent.right
		for node.left != nil {
			node = node.left
		}

		// set the current node value to the least value in the right subtree
		parent.value = node.value
		// remove the found node from its own right subtree
		parent.right = t.deleteFromSubtree(node.value, parent.right)

	}
	return parent
}
