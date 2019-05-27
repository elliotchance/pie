package functions

// Shift will return two values: the shifted value and the rest slice.
func (ss SliceType) Shift() (ElementType, SliceType) {
	return ss.First(), ss.DropTop(1)
}
