package binary_search_tree

import (
	"github.com/AlekSi/pointer"
	"github.com/artemovskiy/golang-data-structures/pkg/iter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Delete(t *testing.T) {
	getTestTree := func() *Tree {
		return &Tree{
			root: &Node{
				value: 5,
				left: &Node{
					value: 3,
					left:  &Node{value: 2},
					right: &Node{value: 4},
				},
				right: &Node{
					value: 7,
					left:  &Node{value: 6},
					right: &Node{value: 8},
				},
			},
		}
	}

	cases := []struct {
		value    int
		expected []*int
	}{
		{
			value: 8,
			expected: []*int{
				pointer.To(5),
				pointer.To(3), pointer.To(7),
				pointer.To(2), pointer.To(4), pointer.To(6),
			},
		},
		{
			value: 6,
			expected: []*int{
				pointer.To(5),
				pointer.To(3), pointer.To(7),
				pointer.To(2), pointer.To(4), nil, pointer.To(8),
			},
		},
		{
			value: 7,
			expected: []*int{
				pointer.To(5),
				pointer.To(3), pointer.To(8),
				pointer.To(2), pointer.To(4), pointer.To(6),
			},
		},
		{
			value: 3,
			expected: []*int{
				pointer.To(5),
				pointer.To(4), pointer.To(7),
				pointer.To(2), nil, pointer.To(6), pointer.To(8),
			},
		},
		{
			value: 5,
			expected: []*int{
				pointer.To(6),
				pointer.To(3), pointer.To(7),
				pointer.To(2), pointer.To(4), nil, pointer.To(8),
			},
		},
	}

	for _, tc := range cases {
		tree := getTestTree()
		tree.Delete(tc.value)
		result := iter.Collect(tree.LevelOrderTraversalWithGaps())
		tree.Dump()

		assert.Equal(t, tc.expected, result)
	}
}
