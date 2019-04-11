package functions

//
func (ss SliceType) Contains(lookingFor ElementType) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}
