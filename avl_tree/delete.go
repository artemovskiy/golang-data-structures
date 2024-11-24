package avl_tree

func (t *Tree) Delete(v int) {
	t.root = t.deleteFromSubtree(v, t.root)
}

func (t *Tree) deleteFromSubtree(v int, parent *Node) *Node {
	if parent == nil {
		return nil
	}
	if v < parent.value {
		parent.left = t.deleteFromSubtree(v, parent.left)
	}
	if v > parent.value {
		parent.right = t.deleteFromSubtree(v, parent.right)
	}
	if v == parent.value {
		if parent.left == nil && parent.right == nil {
			return nil
		} else if parent.left != nil && parent.right != nil {
			node := parent.right
			for node.left != nil {
				node = node.left
			}

			parent.value = node.value

			parent.right = t.deleteFromSubtree(node.value, parent.right)
		} else if parent.left != nil {
			return parent.left
		} else if parent.right != nil {
			return parent.right
		}
	}

	t.fixHeight(parent)
	bf := parent.balanceFactor()

	if bf > 1 {
		if parent.left.balanceFactor() == -1 {
			parent.left = t.rotateLeftOver(parent.left)
		}
		return t.rotateRightOver(parent)
	}

	if bf < -1 {
		if parent.right.balanceFactor() == 1 {
			parent.right = t.rotateRightOver(parent.right)
		}
		return t.rotateLeftOver(parent)
	}

	return parent
}
