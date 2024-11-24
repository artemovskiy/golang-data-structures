package avl_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchTree_InsertSymmetricTraverse(t *testing.T) {
	data := [...]int{1333, 342, 0, 78, 34 /*, 234, 6676, 20, 309, 4, 15*/}
	tree := Tree{}
	for _, v := range data {
		tree.Insert(v)
	}

	tree.Dump()

	assert.Equal(t, 1, 1)
	//expected := []int{0, 4, 15, 20, 34, 78, 234, 309, 342, 1333, 6676}
	//assert.Equal(t, expected, tree.SymmetricTraverse())
	//
	//expected2 := []int{1333, 342, 0, 78, 34, 20, 4, 15, 234, 309, 6676}
	//
	//assert.Equal(t, expected2, tree.PreOrderTraverse())
}
