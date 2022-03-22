package pie

// Map will return a new slice where each element has been mapped (transformed).
// The number of elements returned will always be the same as the input.
//
// Be careful when using this with slices of pointers. If you modify the input
// value it will affect the original slice. Be sure to return a new allocated
// object or deep copy the existing one.
func Map[T any, U any](ss []T, fn func(T) U) (ss2 []U) {
	if ss == nil {
		return nil
	}

	ss2 = make([]U, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}
