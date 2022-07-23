package pie

// SplitSlice returns splited slices which length equals splitLength.
//
// for this [1,2,3] slice with splitLength == 2 will return [][]int{{1,2}, {3}}.
// If input slice length is less than splitLength, for this [1,2,3] slice with splitLength == 4
// 		will return [][]int{{1,2,3}}.
// If input slice length == 0 or splitLength <= 0, will return [][]T{}.
func SplitSlice[T any](ss []T, splitLength int) [][]T {
	result := make([][]T, 0)
	l := len(ss)
	if l == 0 || splitLength <= 0 {
		return result
	}

	var step = l / splitLength
	if step == 0 {
		result = append(result, ss)
		return result
	}
	var remain = l % splitLength
	for i := 0; i < step; i++ {
		result = append(result, ss[i*splitLength:(i+1)*splitLength])
	}
	if remain != 0 {
		result = append(result, ss[step*splitLength:l])
	}
	return result
}
