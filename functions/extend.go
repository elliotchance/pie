package functions

//
func (ss SliceType) Extend(slices ...SliceType) (ss2 SliceType) {
	ss2 = ss

	for _, slice := range slices {
		ss2 = ss2.Append(slice...)
	}

	return ss2
}
