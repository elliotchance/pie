package pie

// ZipLongest will return a new slice containing pairs with elements from input slices.
// If input slices have different length, missing elements will be padded with default values.
func ZipLongest[T1, T2 any](ss1 []T1, ss2 []T2) []Zipped[T1, T2] {
	ss3 := make([]Zipped[T1, T2], Max([]int{len(ss1), len(ss2)}))
	for i := range ss3 {
		ss3[i] = Zipped[T1, T2]{
			A: First(ss1),
			B: First(ss2),
		}
		Pop(&ss1)
		Pop(&ss2)
	}

	return ss3
}
