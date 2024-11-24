package avl_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Find(t *testing.T) {
	cases := []struct {
		data   []int
		value  int
		result bool
	}{
		{
			data:   []int{1, 2, 3},
			value:  1,
			result: true,
		},
		{
			data:   []int{1, 2, 3},
			value:  2,
			result: true,
		},
		{
			data:   []int{1, 2, 3},
			value:  4,
			result: false,
		},
		{
			data:   []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
			value:  8,
			result: true,
		},
		{
			data:   []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
			value:  45,
			result: false,
		},
		{
			data:   []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
			value:  31,
			result: true,
		},
		{
			data:   []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
			value:  6,
			result: false,
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			tree := Tree{}
			for _, insert := range tc.data {
				tree.Insert(insert)
			}

			result := tree.Find(tc.value)

			assert.Equal(t, tc.result, result)
		})
	}
}
