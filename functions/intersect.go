package functions

// Intersect returns items that exist in all lists.
//
// It returns slice without any duplicates.
// If zero slice arguments are provided, then nil is returned.
func (ss SliceType) Intersect(slices ...SliceType) (ss2 SliceType) {
	if slices == nil {
		return nil
	}

	var uniqs []map[ElementType]struct{}
	for _, s := range slices {
		m := make(map[ElementType]struct{})
		for _, el := range s {
			m[el] = struct{}{}
		}
		uniqs = append(uniqs, m)
	}

	var containsInAll = false
	for _, el := range ss.Unique() {
		for _, u := range uniqs {
			if _, exists := u[el]; !exists {
				containsInAll = false
				break
			}
			containsInAll = true
		}
		if containsInAll {
			ss2 = append(ss2, el)
		}
	}

	return
}
