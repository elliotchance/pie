package pie

// IntTransformFunc transforms an int value.
type IntTransformFunc func(int) int

// AddInt is addition.
func AddInt(a int) IntTransformFunc {
	return func(i int) int {
		return a + i
	}
}
