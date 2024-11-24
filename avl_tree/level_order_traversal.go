package avl_tree

import (
	"github.com/AlekSi/pointer"
	"iter"
)

func (t *Tree) LevelOrderTraversal() iter.Seq[*Node] {
	return func(yield func(*Node) bool) {
		queue := make([]*Node, 1)
		queue[0] = t.root
		i := 0
		for i < len(queue) {
			node := queue[i]
			r := yield(node)
			if !r {
				return
			}
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
			i++
		}
	}
}

type nodeLevel struct {
	node   *Node
	level  int
	offset int
}

func (t *Tree) LevelOrderTraversalWithGaps() iter.Seq[*int] {
	return func(yield func(*int) bool) {
		queue := make([]nodeLevel, 1)
		queue[0] = nodeLevel{node: t.root}
		i := 0

		var currentLevel int
		var currentOffset int
		for i < len(queue) {
			nl := queue[i]
			if nl.level > currentLevel {
				levelWidth := 1
				for j := 0; j < currentLevel; j++ {
					levelWidth *= 2
				}
				// fmt.Printf("level %d width %d currentOffset %d \n", currentLevel, levelWidth, currentOffset)
				for currentOffset < (levelWidth) {
					r := yield(nil)
					currentOffset++
					if !r {
						return
					}
				}
				currentLevel = nl.level
				currentOffset = 0
			}
			diff := nl.offset - currentOffset
			if diff > 0 {
				for j := 0; j < diff; j++ {
					r := yield(nil)
					if !r {
						return
					}
				}
				currentOffset += diff
			}

			node := nl.node
			var value *int
			if node != nil {
				value = pointer.To(node.value)
			}
			r := yield(value)
			currentOffset += 1
			if !r {
				return
			}
			if node != nil {
				if node.left != nil {
					queue = append(queue, nodeLevel{node: node.left, level: nl.level + 1, offset: nl.offset * 2})
				}

				if node.right != nil {
					queue = append(queue, nodeLevel{node: node.right, level: nl.level + 1, offset: nl.offset*2 + 1})
				}
			}
			i++
		}

	}
}
