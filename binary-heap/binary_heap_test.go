package binary_heap

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func testHeap(elements []int, i int) bool {
	if i >= len(elements) {
		return true
	}
	l := 2*i + 1
	r := 2*i + 2

	result := true
	if l < len(elements) {
		result = elements[l] < elements[i] && testHeap(elements, l)
	}
	if r < len(elements) {
		result = result && elements[r] < elements[i] && testHeap(elements, r)
	}

	return result
}

func TestBinaryHeap_Insert(t *testing.T) {
	cases := map[string]struct {
		inserts []int
	}{
		"inserting ordered sequence": {
			inserts: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		"inserting reverse - ordered sequence": {
			inserts: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			heap := BinaryHeap{}
			for i := 0; i < len(tc.inserts); i++ {
				heap.Insert(tc.inserts[i])
			}

			require.True(t, testHeap(heap.elements, 0))
			require.Len(t, heap.elements, len(tc.inserts))
		})
	}
}

func TestBinaryHeap_heapify(t *testing.T) {
	cases := map[string]struct {
		elements []int
		expected []int
	}{
		"putting root node to the right place": {
			elements: []int{8, 15, 11, 6, 9, 7, 8, 1, 3, 5},
			expected: []int{15, 9, 11, 6, 8, 7, 8, 1, 3, 5},
		},
		"heap is already right - do nothing": {
			elements: []int{15, 9, 11, 6, 8, 7, 8, 1, 3, 5},
			expected: []int{15, 9, 11, 6, 8, 7, 8, 1, 3, 5},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			heap := BinaryHeap{
				elements: tc.elements,
			}

			heap.heapify(0)

			require.Equal(t, tc.expected, heap.elements)
		})
	}
}

func TestBinaryHeap_GetMax(t *testing.T) {
	cases := map[string]struct {
		elements     []int
		expectedHeap []int
		expectedMax  int
	}{
		"get max and then heapify": {
			elements:     []int{15, 9, 11, 6, 8, 7, 8, 1, 3, 5},
			expectedHeap: []int{11, 9, 8, 6, 8, 7, 5, 1, 3},
			expectedMax:  15,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			heap := BinaryHeap{
				elements: tc.elements,
			}

			maxEl := heap.GetMax()

			require.Equal(t, tc.expectedMax, maxEl)
			require.Equal(t, tc.expectedHeap, heap.elements)
		})
	}
}

func TestBinaryHeap_Sorted(t *testing.T) {
	cases := map[string]struct {
		elements []int
		expected []int
	}{
		"get max and then heapify": {
			elements: []int{15, 9, 11, 6, 8, 7, 8, 1, 3, 5},
			expected: []int{1, 3, 5, 6, 7, 8, 8, 9, 11, 15},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			heap := BinaryHeap{
				elements: tc.elements,
			}

			require.Equal(t, tc.expected, heap.Sorted())
		})
	}
}
