package functions

// Mode returns a new slice containing the most frequently occuring values.
//
// The number of items returned may be the same as the input or less. It will
// never return zero items unless then input slice has zero items.
func (ss SliceType) Mode() SliceType {
	values := make(map[ElementType]int)
	for _, s:= range ss {
		values[s]++
	}

	var maxFrequency int
	
	for _,v := range values {
		if v > maxFrequency {
			maxFrequency = v
		}
	}
	
	var maxValues SliceType
	for k,v := range values {
		if v == maxFrequency {
			maxValues = append(maxValues, k)
		}
	}

	return maxValues 
}
