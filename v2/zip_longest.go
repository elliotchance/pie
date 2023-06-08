package pie

// ZipLongest will return a new slice containing pairs with elements from input slices.
// If input slices have diffrent length, missing elements will be padded with default values.
func ZipLongest[T1, T2 any](ss1 []T1, ss2 []T2) []Zipped[T1, T2] {
	var minLen, maxLen int
	var small int8

	if len(ss1) <= len(ss2) {
		small = 1
		minLen = len(ss1)
		maxLen = len(ss2)
	} else {
		small = 2
		minLen = len(ss2)
		maxLen = len(ss1)
	}

	ss3 := []Zipped[T1, T2]{}
	for i := 0; i < minLen; i++ {
		ss3 = append(ss3, Zipped[T1, T2]{ss1[i], ss2[i]})
	}

	if small == 1 {
		var t T1
		for i := minLen; i < maxLen; i++ {
			ss3 = append(ss3, Zipped[T1, T2]{t, ss2[i]})
		}
	} else {
		var t T2
		for i := minLen; i < maxLen; i++ {
			ss3 = append(ss3, Zipped[T1, T2]{ss1[i], t})
		}
	}

	return ss3
}
