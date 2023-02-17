package pie

// Flat flattens the two-dimensional slice into one-dimensional slice.
// Slices of zero-length are ignored.
//
// Examples:
//
//	Flat([[100], [101, 102], [102, 103]])   => [100, 101, 102, 102, 103]
//	Flat([nil, [101, 102], []])             => [101, 102]
func Flat[T any](ss [][]T) (ss2 []T) {
	for _, s := range ss {
		ss2 = append(ss2, s...)
	}

	return ss2
}
