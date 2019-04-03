package pie

// AddFloat64 is addition.
func AddFloat64(a float64) Float64sTransformFunc {
	return func(i float64) float64 {
		return a + i
	}
}

// EqualFloat64 return true if the float64s are equal.
func EqualFloat64(a float64) Float64sConditionFunc {
	return func(i float64) bool {
		return a == i
	}
}
