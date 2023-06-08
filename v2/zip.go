package pie

// A pair struct containing two zipped values.
type Zipped[T1, T2 any] struct {
	A T1
	B T2
}

// Zip will return a new slice containing pairs with elements from input slices.
// If input slices have diffrent length, the output slice will be truncated to
// the length of the smallest input slice.
func Zip[T1, T2 any](ss1 []T1, ss2 []T2) []Zipped[T1, T2] {
	var minLen int

	if len(ss1) <= len(ss2) {
		minLen = len(ss1)
	} else {
		minLen = len(ss2)
	}

	ss3 := []Zipped[T1, T2]{}
	for i := 0; i < minLen; i++ {
		ss3 = append(ss3, Zipped[T1, T2]{ss1[i], ss2[i]})
	}

	return ss3
}
