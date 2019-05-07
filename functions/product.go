package functions

// Product is the product of all of the elements.
func (ss SliceType) Product() (product ElementType) {
	if len(ss) == 0 {
		return
	}
	product = ss[0]
	for _, s := range ss[1:] {
		product *= s
	}

	return
}
