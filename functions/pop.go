package functions

//Pop is method get one element at the end of slices, or zero value
func (ss SliceType) Pop() ElementType {
	if len(ss) == 0 {
		return
	}
	popValue, ss = ss[len(ss)-1], ss[:len(ss)-1]

	return popValue
}
