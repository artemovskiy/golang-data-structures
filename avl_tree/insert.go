package avl_tree

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func (t *Tree) Insert(value int) {
	if t.root == nil {
		t.root = &Node{value: value, height: 1}
	} else {
		t.insertWithParent(value, t.root)
	}
}

func (t *Tree) insertWithParent(value int, parent *Node) *Node {
	if parent == nil {
		return &Node{value: value, height: 1}
	}
	if value < parent.value {
		parent.left = t.insertWithParent(value, parent.left)
	} else if value > parent.value {
		parent.right = t.insertWithParent(value, parent.right)
	} else {
		return parent
	}

	t.fixHeight(parent)
	bf := height(parent.left) - height(parent.right)
	if bf > 1 {
		if parent.left.balanceFactor() == -1 {
			parent.left = t.rotateLeftOver(parent.left)
		}
		return t.rotateRightOver(parent)
	}

	if bf < -1 {
		if parent.right.balanceFactor() == 1 {
			parent.right = t.rotateRightOver(parent.right)
			return t.rotateLeftOver(parent)
		} else {
			return t.rotateLeftOver(parent)
		}
	}

	return parent
}

func (t *Tree) rotateRightOver(node *Node) *Node {
	// originalParent := node.parent
	subtreeRoot := node.left
	node.left = subtreeRoot.right
	subtreeRoot.right = node

	t.fixHeight(node)
	t.fixHeight(subtreeRoot)

	if node == t.root {
		t.root = subtreeRoot
	}

	return subtreeRoot

}

func (t *Tree) rotateLeftOver(a *Node) *Node {
	// originalParent := a.parent
	b := a.right
	a.right = b.left
	b.left = a

	t.fixHeight(a)
	t.fixHeight(b)

	if a == t.root {
		t.root = b
	}

	return b

}

func (t *Tree) fixHeight(node *Node) {
	node.height = 1 + max(height(node.left), height(node.right))
}

func (n *Node) balanceFactor() int {
	return height(n.left) - height(n.right)
}
