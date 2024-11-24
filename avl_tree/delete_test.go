package avl_tree

import (
	"github.com/AlekSi/pointer"
	"github.com/artemovskiy/golang-data-structures/pkg/iter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Delete(t *testing.T) {
	getTestTree := func() *Tree {
		tree := Tree{}
		for i := 0; i < 13; i++ {
			tree.Insert(i)
		}
		return &tree
	}

	cases := map[string]struct {
		deletes  []int
		expected []*int
	}{
		"delete a leaf 0 on the 3rd level": {
			deletes: []int{0},
			expected: []*int{
				pointer.To(7),
				pointer.To(3), pointer.To(9),
				pointer.To(1), pointer.To(5), pointer.To(8), pointer.To(11),
				nil, pointer.To(2), pointer.To(4), pointer.To(6), nil, nil, pointer.To(10), pointer.To(12),
			},
		},
		"delete a leaf 8 on the 2nd level (w l rotation)": {
			deletes: []int{8},
			expected: []*int{
				pointer.To(7),
				pointer.To(3), pointer.To(11),
				pointer.To(1), pointer.To(5), pointer.To(9), pointer.To(12),
				pointer.To(0), pointer.To(2), pointer.To(4), pointer.To(6), nil, pointer.To(10),
			},
		},
		"delete node 1 on 2nd level(wo rotation)": {
			deletes: []int{1},
			expected: []*int{
				pointer.To(7),
				pointer.To(3), pointer.To(9),
				pointer.To(2), pointer.To(5), pointer.To(8), pointer.To(11),
				pointer.To(0), nil, pointer.To(4), pointer.To(6), nil, nil, pointer.To(10), pointer.To(12),
			},
		},
		"delete root node 7 (w left rotation)": {
			deletes: []int{7},
			expected: []*int{
				pointer.To(8),
				pointer.To(3), pointer.To(11),
				pointer.To(1), pointer.To(5), pointer.To(9), pointer.To(12),
				pointer.To(0), pointer.To(2), pointer.To(4), pointer.To(6), nil, pointer.To(10),
			},
		},
		"delete 0,4,6,5(big right rotation)": {
			deletes: []int{0, 4, 6, 5},
			expected: []*int{
				pointer.To(7),
				pointer.To(2), pointer.To(9),
				pointer.To(1), pointer.To(3), pointer.To(8), pointer.To(11),
				nil, nil, nil, nil, nil, nil, pointer.To(10), pointer.To(12),
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			tree := getTestTree()
			for _, v := range tc.deletes {
				tree.Delete(v)
			}
			result := iter.Collect(tree.LevelOrderTraversalWithGaps())
			tree.Dump()

			assert.Equal(t, tc.expected, result)
		})
	}
}
