package pie

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss myInt64s) Average() float64 {
	if l := int64(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}

// Sum is the sum of all of the elements.
func (ss myInt64s) Sum() (sum int64) {
	for _, s := range ss {
		sum += s
	}

	return
}
