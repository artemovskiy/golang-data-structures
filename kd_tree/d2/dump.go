package d2

import "fmt"

func (t *Tree) Dump() {
	t.dumpFromNode(t.root, []int{})
	fmt.Print("\n")
}

func (t *Tree) dumpFromNode(node *Node, y []int) {
	if node == nil {
		return
	}
	var nodeStr string
	if node.dimension == DimensionNs {
		nodeStr = fmt.Sprintf("[%d:%d]", node.point.x, node.point.y)
	} else {
		nodeStr = fmt.Sprintf("[%s:%d]", node.dimension, node.value)
	}
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
