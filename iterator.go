package binary_trees

type Item int

func iterateItems(yield func(Item) bool) {
	items := []Item{1, 2, 3}
	for _, v := range items {
		if !yield(v) {
			return
		}
	}
}
