package pie

// Rotate return slice circularly rotated by a number of positions n.
// If n is positive, the slice is rotated right.
// If n is negative, the slice is rotated left.
func Rotate[T any](ss []T, n int) []T {

	length := len(ss)

	// Avoid the allocation.
	// If there is one element or less, then already rotated.
	if length < 2 {
		return ss
	}

	// Normalize shift
	// no div by 0 since length >= 2
	shift := -n % length
	if shift < 0 {
		shift = length + shift
	}

	// Avoid the allocation.
	// If normalized shift is 0, then already rotated.
	if shift == 0 {
		return ss
	}

	return append(DropTop(ss, shift), Top(ss, shift)...)
}
