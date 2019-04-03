package pie

// IntConditionFunc allows ints to be filtered or checked by value.
type IntConditionFunc func(int) bool

// EqualInt return true if the ints are equal.
func EqualInt(b int) IntConditionFunc {
	return func(a int) bool {
		return a == b
	}
}

// NotEqualInt returns true if the integers exactly equal.
func NotEqualInt(b int) IntConditionFunc {
	return func(a int) bool {
		return a != b
	}
}

// GreaterThanInt returns true if the value is greater than b.
func GreaterThanInt(b int) IntConditionFunc {
	return func(a int) bool {
		return a > b
	}
}

// GreaterThanEqualInt returns true if the value is greater than or equal to b.
func GreaterThanEqualInt(b int) IntConditionFunc {
	return func(a int) bool {
		return a >= b
	}
}

// LessThanInt returns true if the value is greater than b.
func LessThanInt(b int) IntConditionFunc {
	return func(a int) bool {
		return a < b
	}
}

// LessThanEqualInt returns true if the value is greater than or equal to b.
func LessThanEqualInt(b int) IntConditionFunc {
	return func(a int) bool {
		return a <= b
	}
}
