package functions

//
func (ss SliceType) Average() float64 {
	if l := ElementType(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}
