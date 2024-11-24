package binary_trees

import (
	"fmt"
	"iter"
)

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	value  int
}

type BinarySearchTree struct {
	root *Node
}

func (t *BinarySearchTree) Insert(value int) {
	if t.root == nil {
		t.root = &Node{value: value}
	} else {
		t.insertWithParent(value, t.root)
	}
}

// insertWithParent returns a node if the value is already in within the tree
func (t *BinarySearchTree) insertWithParent(value int, parent *Node) *Node {
	if parent.value == value {
		return parent
	}
	if value < parent.value {
		if parent.left == nil {
			parent.left = &Node{value: value, parent: parent}
		} else {
			return t.insertWithParent(value, parent.left)
		}
	} else {
		if parent.right == nil {
			parent.right = &Node{value: value, parent: parent}
		} else {
			return t.insertWithParent(value, parent.right)
		}
	}
	return nil
}

func (t *BinarySearchTree) SymmetricTraverse() []int {
	if t.root == nil {
		return []int{}
	} else {
		return t.symmetricTraverseFrom(t.root)
	}
}

func (t *BinarySearchTree) SymmetricTraverseSeq() iter.Seq[int] {
	if t.root == nil {
		return func(yield func(int) bool) {
			return
		}
	} else {
		return func(yield func(int) bool) {
			t.symmetricTraverseYieldFrom(t.root, yield)
			return
		}
	}
}

func (t *BinarySearchTree) symmetricTraverseYieldFrom(node *Node, yield func(int) bool) bool {
	var result bool
	if node.left != nil {
		result = t.symmetricTraverseYieldFrom(node.left, yield)
		if !result {
			return false
		}
	}
	result = yield(node.value)
	if !result {
		return false
	}
	if node.right != nil {
		return t.symmetricTraverseYieldFrom(node.right, yield)
	}
	return true
}

func (t *BinarySearchTree) symmetricTraverseFrom(node *Node) []int {

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

func (t *BinarySearchTree) PreOrderTraverse() []int {
	if t.root == nil {
		return []int{}
	} else {
		return t.preOrderTraverseFrom(t.root)
	}
}

func (t *BinarySearchTree) preOrderTraverseFrom(node *Node) []int {
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

func (t *BinarySearchTree) Dump() {
	t.dumpFromNode(t.root, []int{})
	fmt.Print("\n")
}

func (t *BinarySearchTree) dumpFromNode(node *Node, y []int) {
	nodeStr := fmt.Sprintf("[%d]", node.value)
	fmt.Print(nodeStr)
	if node.left != nil {
		fmt.Print("─")
		st := make([]int, len(y))
		copy(st, y)
		if node.right != nil {
			st = append(st, 1)
		} else {
			st = append(st, 0)
		}
		for i := 0; i < (len(nodeStr)); i++ {
			st = append(st, 0)
		}
		t.dumpFromNode(node.left, st)
	}

	if node.right != nil {
		fmt.Print("\n")
		for i := 0; i < len(y); i++ {
			if y[i] > 0 {
				fmt.Print("│")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Printf("%c", '└')

		t.dumpFromNode(node.right, append(y, 0))
	}
}
