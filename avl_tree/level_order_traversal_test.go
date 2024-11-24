package avl_tree

import (
	"github.com/AlekSi/pointer"
	"github.com/artemovskiy/golang-data-structures/pkg/iter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchTree_LevelOrderTraversal(t *testing.T) {
	node := Node{
		value: 342,
		right: &Node{value: 1333},
		left: &Node{
			value: 34,
			left:  &Node{value: 0},
			right: &Node{value: 78},
		},
	}
	tree := Tree{root: &node}

	var result []int
	for n := range tree.LevelOrderTraversal() {
		result = append(result, n.value)
	}

	expected := []int{342, 34, 1333, 0, 78}
	assert.Equal(t, expected, result)
}

func TestBinarySearchTree_LevelOrderTraversalWithGaps(t *testing.T) {
	node := Node{
		value: 342,
		left: &Node{
			value: 34,
			left: &Node{
				value: 0,
				right: &Node{
					value: 4,
					left:  &Node{value: 3},
					right: &Node{
						value: 5,
						right: &Node{value: 6},
					},
				},
			},
			right: &Node{value: 78},
		},
		right: &Node{
			value: 1333,
			right: &Node{value: 1334},
		},
	}
	tree := Tree{root: &node}

	result := iter.Collect(tree.LevelOrderTraversalWithGaps())

	expected := []*int{
		pointer.To(342),
		pointer.To(34), pointer.To(1333),
		pointer.To(0), pointer.To(78), nil, pointer.To(1334),
		nil, pointer.To(4), nil, nil, nil, nil, nil, nil,
		nil, nil, pointer.To(3), pointer.To(5), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, pointer.To(6),
	}
	assert.Equal(t, expected, result)
}
