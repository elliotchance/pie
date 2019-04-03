package pie

// AddInt is addition.
func AddInt(a int) IntsTransformFunc {
	return func(i int) int {
		return a + i
	}
}

// EqualInt return true if the ints are equal.
func EqualInt(a int) IntsConditionFunc {
	return func(i int) bool {
		return a == i
	}
}
