package pie

import "strings"

// StringConditionFunc allows strings to be filtered or checked by value.
type StringConditionFunc func(string) bool

// Prefix is a wrapper for strings.HasPrefix that can be used with any function
// that accepts a StringConditionFunc. For example:
//
//   namesStartingWithJ := pie.Strings{"Bob", "Sally", "John", "Jane"}.
//   	Only(func(s string) bool {
//   		return len(s) > 0 && s[0] == 'J'
//   	})
//
// Could more easily be written as:
//
//   namesStartingWithJ := pie.Strings{"Bob", "Sally", "John", "Jane"}.
//   	Only(pie.Prefix("J"))
//
func Prefix(s string) StringConditionFunc {
	return func(a string) bool {
		return strings.HasPrefix(a, s)
	}
}

// Suffix is a wrapper for strings.HasSuffix. See HasPrefix for documentation.
func Suffix(s string) StringConditionFunc {
	return func(a string) bool {
		return strings.HasSuffix(a, s)
	}
}

// EqualString returns true if the strings are exactly equal (case-sensitive).
func EqualString(b string) StringConditionFunc {
	return func(a string) bool {
		return a == b
	}
}

// NotEqualString returns true if the strings are exactly equal
// (case-sensitive).
func NotEqualString(b string) StringConditionFunc {
	return func(a string) bool {
		return a != b
	}
}

// GreaterThanString returns true if the value is greater than s.
func GreaterThanString(b string) StringConditionFunc {
	return func(a string) bool {
		return a > b
	}
}

// GreaterThanEqualString returns true if the value is greater than or equal to
// s.
func GreaterThanEqualString(b string) StringConditionFunc {
	return func(a string) bool {
		return a >= b
	}
}

// LessThanString returns true if the value is greater than s.
func LessThanString(b string) StringConditionFunc {
	return func(a string) bool {
		return a < b
	}
}

// LessThanEqualString returns true if the value is greater than or equal to s.
func LessThanEqualString(b string) StringConditionFunc {
	return func(a string) bool {
		return a <= b
	}
}
