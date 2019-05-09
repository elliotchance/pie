package functions

//Pop wil return a new slice and with poped value which element at the end of slice or return nil and 0 (or "" if Strings) if slice empty
func (ss SliceType) Pop() (SliceType, ElementType) {
	if len(ss) == 0 {
		return nil, ElementZeroValue
	}
	popValue := ss[len(ss)-1]
	ss = ss[:len(ss)-1]
	return ss, popValue
}
