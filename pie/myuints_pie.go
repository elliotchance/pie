package pie

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss myUints) Average() float64 {
	if l := uint(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}

// Sum is the sum of all of the elements.
func (ss myUints) Sum() (sum uint) {
	for _, s := range ss {
		sum += s
	}

	return
}
