package functions

// IntersectUsing returns items that exist in all lists, using equals function
//
// It returns slice without any duplicates.
// If zero slice arguments are provided, then nil is returned.
func (ss SliceType) IntersectUsing(equals func(ElementType, ElementType) bool, slices ...SliceType) (ss2 SliceType) {
	if slices == nil {
		return nil
	}

	for _, e1 := range ss {
		for _, s2 := range slices {
			for _, e2 := range s2 {
				if equals(e1, e2) {
					ss2 = append(ss2, e1)
				}
			}
		}
	}

	return ss2.Unique()
}
