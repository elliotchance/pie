package functions

//Push is method pass one element as argument and add element to the end of slice, or zero value
func (ss SliceType) Push(element ElementType) SliceType {
	ss = append(ss, element)
	return ss
}
