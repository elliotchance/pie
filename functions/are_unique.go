package functions

//
func (ss SliceType) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
}
