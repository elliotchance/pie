package pie

// Float64ConditionFunc allows float64s to be filtered or checked by value.
type Float64ConditionFunc func(float64) bool

// EqualFloat64 return true if the ints are equal.
func EqualFloat64(b float64) Float64ConditionFunc {
	return func(a float64) bool {
		return a == b
	}
}

// NotEqualFloat64 returns true if the integers exactly equal.
func NotEqualFloat64(b float64) Float64ConditionFunc {
	return func(a float64) bool {
		return a != b
	}
}

// GreaterThanFloat64 returns true if the value is greater than b.
func GreaterThanFloat64(b float64) Float64ConditionFunc {
	return func(a float64) bool {
		return a > b
	}
}

// GreaterThanEqualFloat64 returns true if the value is greater than or equal to b.
func GreaterThanEqualFloat64(b float64) Float64ConditionFunc {
	return func(a float64) bool {
		return a >= b
	}
}

// LessThanFloat64 returns true if the value is greater than b.
func LessThanFloat64(b float64) Float64ConditionFunc {
	return func(a float64) bool {
		return a < b
	}
}

// LessThanEqualFloat64 returns true if the value is greater than or equal to b.
func LessThanEqualFloat64(b float64) Float64ConditionFunc {
	return func(a float64) bool {
		return a <= b
	}
}
