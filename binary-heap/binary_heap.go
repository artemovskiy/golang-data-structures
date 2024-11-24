package binary_heap

type BinaryHeap struct {
	elements []int
}

func (h *BinaryHeap) Insert(v int) {
	h.elements = append(h.elements, v)
	i := len(h.elements) - 1
	// if ith element is a left child, its index is 2p+1, so 2p = i - 1 => p = (i - 1) / 2
	// if ith element is a right child, its index is 2p+2, so 2p = i - 2 => 2p + 1 = i - 1 => 2p + 1 mod 2 = i - 1 => p = (i - 1) / 2 | mod 1
	parent := (i - 1) / 2
	for parent >= 0 && h.elements[parent] < h.elements[i] {
		h.elements[parent], h.elements[i] = h.elements[i], h.elements[parent]
		i = parent
		parent = (parent - 1) / 2
	}
}

func (h *BinaryHeap) GetMax() int {
	maxEl := h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	if len(h.elements) > 0 {
		h.heapify(0)
	}
	return maxEl
}

func (h *BinaryHeap) heapify(i int) {
	for {
		v := h.elements[i]
		l := 2*i + 1
		r := 2*i + 2

		lv := v - 1
		if l < len(h.elements) {
			lv = h.elements[l]
		}

		rv := v - 1
		if r < len(h.elements) {
			rv = h.elements[r]
		}

		if lv <= v && rv <= v {
			break
		}

		if rv > lv {
			h.elements[i], h.elements[r] = rv, v
			i = r
		} else {
			h.elements[i], h.elements[l] = lv, v
			i = l
		}
	}
}

func (h *BinaryHeap) Sorted() []int {
	arr := make([]int, len(h.elements))
	for i := len(h.elements) - 1; i >= 0; i-- {
		arr[i] = h.GetMax()
	}
	return arr
}

func BuildHeap(elements []int) *BinaryHeap {
	h := BinaryHeap{elements: elements}
	for i := len(elements) / 2; i >= 0; i-- {
		h.heapify(i)
	}
	return &h
}
