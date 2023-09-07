package pie

// Shift will return two values: the shifted value and the rest slice.
// if the slice is empty then returned shifted value is the zero value of the slice elements and the rest slice is empty slice
func Shift[T any](ss []T) (T, []T) {
	var zeroValue T
	return FirstOr(ss, zeroValue), DropTop(ss, 1)
}
