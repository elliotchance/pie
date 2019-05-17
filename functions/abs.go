package functions

// Abs is a function which returns the absolute value of all the
// elements in the slice.
func (ss SliceType) Abs() SliceType {
	result := make(SliceType, len(ss))
	for i, val := range ss {
		if val < 0 {
			result[i] = -val
		} else {
			result[i] = val
		}
	}
	return result
}
