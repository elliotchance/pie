package pie

// Reduce continually applies the provided function
// over the slice. Reducing the elements to a single value.
//
// Returns a zero value of T if there are no elements in the slice. It will
// panic if the reducer is nil and the slice has more than one element (required
// to invoke reduce). Otherwise returns result of applying reducer from left to
// right.
func Reduce[T any](ss []T, reducer func(T, T) T) (el T) {
	if len(ss) == 0 {
		return
	}

	el = ss[0]
	for _, s := range ss[1:] {
		el = reducer(el, s)
	}

	return
}
