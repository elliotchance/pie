package functions

// Drop items from the slice while f(item) is true.
// Afterwards, return every element until the slice is empty. It follows the same logic as the dropwhile() function from itertools in Python.
func (ss SliceType) DropWhile(f func(s ElementType) bool) (ss2 SliceType) {
	ss2 = make([]ElementType, len(ss))
	copy(ss2, ss)
	for i, value := range ss2 {
		if !f(value) {
			return ss2[i:]
		}
	}
	return SliceType{}
}
