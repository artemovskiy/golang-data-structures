package avl_tree

import (
	"fmt"
	"slices"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/artemovskiy/golang-data-structures/pkg/iter"
	"github.com/stretchr/testify/assert"
)

func reverseSlice(s []int) []int {
	r := make([]int, len(s))
	copy(r, s)
	slices.Reverse(r)
	return r
}

func TestTree_Insert(t *testing.T) {
	cases := []struct {
		inserts  []int
		expected []*int
	}{
		{
			inserts: []int{1, 2, 3},
			expected: []*int{
				pointer.ToInt(2),
				pointer.ToInt(1), pointer.ToInt(3),
			},
		},
		{
			inserts: []int{3, 2, 1},
			expected: []*int{
				pointer.ToInt(2),
				pointer.ToInt(1), pointer.ToInt(3),
			},
		},
		{
			inserts: []int{8, 6, 4, 10, 9},
			expected: []*int{
				pointer.ToInt(6),
				pointer.ToInt(4), pointer.ToInt(9),
				nil, nil, pointer.ToInt(8), pointer.ToInt(10),
			},
		},
		{
			inserts: []int{8, 6, 5, 2, 4, 3},
			expected: []*int{
				pointer.ToInt(4),
				pointer.ToInt(2), pointer.ToInt(6),
				nil, pointer.To(3), pointer.To(5), pointer.To(8),
			},
		},
		{
			inserts: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
			expected: []*int{
				// level 0
				pointer.ToInt(16),
				// level 1
				pointer.ToInt(8), pointer.ToInt(24),
				// level 2
				pointer.To(4), pointer.To(12), pointer.To(20), pointer.To(28),
				// level 3
				pointer.To(2), pointer.To(6), pointer.To(10), pointer.To(14),
				pointer.To(18), pointer.To(22), pointer.To(26), pointer.To(30),
				// level 4
				pointer.To(1), pointer.To(3), pointer.To(5), pointer.To(7),
				pointer.To(9), pointer.To(11), pointer.To(13), pointer.To(15),
				pointer.To(17), pointer.To(19), pointer.To(21), pointer.To(23),
				pointer.To(25), pointer.To(27), pointer.To(29), pointer.To(31),
			},
		},
		{
			inserts: reverseSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}),
			expected: []*int{
				// level 0
				pointer.ToInt(16),
				// level 1
				pointer.ToInt(8), pointer.ToInt(24),
				// level 2
				pointer.To(4), pointer.To(12), pointer.To(20), pointer.To(28),
				// level 3
				pointer.To(2), pointer.To(6), pointer.To(10), pointer.To(14),
				pointer.To(18), pointer.To(22), pointer.To(26), pointer.To(30),
				// level 4
				pointer.To(1), pointer.To(3), pointer.To(5), pointer.To(7),
				pointer.To(9), pointer.To(11), pointer.To(13), pointer.To(15),
				pointer.To(17), pointer.To(19), pointer.To(21), pointer.To(23),
				pointer.To(25), pointer.To(27), pointer.To(29), pointer.To(31),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			tree := Tree{}
			for _, insert := range tc.inserts {
				fmt.Printf("before insert %d \n", insert)
				tree.Dump()
				tree.Insert(insert)
			}

			result := iter.Collect(tree.LevelOrderTraversalWithGaps())
			tree.Dump()

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestTree_rotateRightOverNode(t *testing.T) {
	tree := Tree{
		root: &Node{
			value: 4,
			left: &Node{
				value: 2,
				left: &Node{
					value: 1,
				},
				right: &Node{
					value: 3,
				},
			},
		},
	}

	tree.rotateRightOver(tree.root)

	expected := []*int{
		pointer.To(2),
		pointer.To(1), pointer.To(4),
		nil, nil, pointer.To(3),
	}
	assert.Equal(t, expected, iter.Collect(tree.LevelOrderTraversalWithGaps()))
}

func TestTree_rotateLeftOverNode(t *testing.T) {
	tree := Tree{
		root: &Node{
			value: 2,
			right: &Node{
				value: 6,
				left: &Node{
					value: 4,
				},
				right: &Node{
					value: 7,
				},
			},
		},
	}

	tree.rotateLeftOver(tree.root)

	tree.Dump()

	expected := []*int{
		pointer.To(6),
		pointer.To(2), pointer.To(7),
		nil, pointer.To(4),
	}
	assert.Equal(t, expected, iter.Collect(tree.LevelOrderTraversalWithGaps()))
}
