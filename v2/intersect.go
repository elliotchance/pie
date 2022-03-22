package pie

// Intersect returns items that exist in all lists.
//
// It returns slice without any duplicates.
// If zero slice arguments are provided, then nil is returned.
func Intersect[T comparable](ss []T, slices ...[]T) (ss2 []T) {
	if slices == nil {
		return nil
	}

	var uniqs = make([]map[T]struct{}, len(slices))
	for i := 0; i < len(slices); i++ {
		m := make(map[T]struct{})
		for _, el := range slices[i] {
			m[el] = struct{}{}
		}
		uniqs[i] = m
	}

	var containsInAll = false
	for _, el := range Unique(ss) {
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
