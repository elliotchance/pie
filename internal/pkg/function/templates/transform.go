package templates

// Transform will return a new slice where each element has been transformed.
// The number of element returned will always be the same as the input.
//
// Be careful when using this with slices of pointers. If you modify the input
// value it will affect the original slice. Be sure to return a new allocated
// object or deep copy the existing one.
func (ss SliceType) Transform(fn func(ElementType) ElementType) (ss2 SliceType) {
	if ss == nil {
		return nil
	}

	ss2 = make([]ElementType, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}
