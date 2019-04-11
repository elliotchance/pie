package main

// The functions in this file only work on numeric slices.

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss SliceType) Average() float64 {
	if l := ElementType(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}

// Sum is the sum of all of the elements.
func (ss SliceType) Sum() (sum ElementType) {
	for _, s := range ss {
		sum += s
	}

	return
}
