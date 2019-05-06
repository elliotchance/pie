package functions

// Reduce continually applies the provided function
// over the slice. Reducing the elements to a single value.
//
// returns: zero value of type if no elements in slice or no function provided.
// Otherwise returns result of applying reducer from left to right.
func (ss SliceType) Reduce(reducer func(ElementType, ElementType) ElementType) (el ElementType) {
	len(ss) == 0 {
		return
	}
	el = ss[0]
	for _, s := range ss[1:] {
		el = reducer(el, s)
	}
	return
}
