package functions

// Reduce continually applies the provided function
// over the slice. Reducing the elements to a single value.
//
// returns: nil if no elements in slice, or result of applying reducer from left to right.
func (ss SliceType) Reduce(reducer func(ElementType, ElementType) ElementType) (el ElementType) {
	if reducer == nil || len (ss) == 0{
		return 
	}
	el = ss[0]
	for _, s := range ss[1:] {
	    el = reducer(el, s)
	}
	return
}
