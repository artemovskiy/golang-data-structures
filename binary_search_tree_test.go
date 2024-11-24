package binary_trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchTree_InsertSymmetricTraverse(t *testing.T) {
	data := [...]int{1333, 342, 0, 78, 34, 234, 6676, 20, 309, 4, 15}
	tree := BinarySearchTree{}
	for _, v := range data {
		tree.Insert(v)
	}

	expected := []int{0, 4, 15, 20, 34, 78, 234, 309, 342, 1333, 6676}

	tree.Dump()
	assert.Equal(t, expected, tree.SymmetricTraverse())

	expected2 := []int{1333, 342, 0, 78, 34, 20, 4, 15, 234, 309, 6676}

	assert.Equal(t, expected2, tree.PreOrderTraverse())
}

func TestBinarySearchTree_InsertSymmetricTraverseByIterator(t *testing.T) {
	data := [...]int{1333, 342, 0, 78, 34, 234, 6676, 20, 309, 4, 15}
	tree := BinarySearchTree{}
	for _, v := range data {
		tree.Insert(v)
	}

	expected := []int{0, 4, 15, 20, 34, 78, 234, 309, 342, 1333, 6676}

	i := 0
	for v := range tree.SymmetricTraverseSeq() {
		assert.Equal(t, expected[i], v)
		i++

	}
}

func TestBinarySearchTree_InsertSymmetricTraverseByIteratorPartial(t *testing.T) {
	data := [...]int{1333, 342, 0, 78, 34, 234, 6676, 20, 309, 4, 15}
	tree := BinarySearchTree{}
	for _, v := range data {
		tree.Insert(v)
	}

	expected := []int{0, 4, 15, 20, 34, 78, 234, 309, 342, 1333, 6676}

	i := 0
	for v := range tree.SymmetricTraverseSeq() {
		assert.Equal(t, expected[i], v)
		i++

		if i > len(expected)/2 {
			break
		}
	}

	assert.Equal(t, 6, i)
}
