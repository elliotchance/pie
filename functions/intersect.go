package functions

// Intersect returns items that exist in all lists.
//
// if there are no this kind of items it will return nil
func (ss SliceType) Intersect(slices ...SliceType) (ss2 SliceType) {
	if slices == nil {
		return nil
	}

	var uniqs []SliceType
	for _, s := range slices {
		uniqs = append(uniqs, s.Unique())
	}

	containsInAll := false
	for _, el := range ss.Unique() {
		for _, u := range uniqs {
			if !u.Contains(el) {
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
