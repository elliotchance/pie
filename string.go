package pie

import "strings"

// HasPrefix is a wrapper for strings.HasPrefix that can be used with any
// function that accepts as StringsConditionFunc. For example:
//
//   namesStartingWithJ := pie.Strings{"Bob", "Sally", "John", "Jane"}.
//   	If(func(s string) bool {
//   		return s[0] == 'J'
//   	})
//
// Could more easily be written as:
//
//   namesStartingWithJ := pie.Strings{"Bob", "Sally", "John", "Jane"}.
//   	If(pie.HasPrefix("J"))
//
func HasPrefix(s string) StringsConditionFunc {
	return func(a string) bool {
		return strings.HasPrefix(a, s)
	}
}

// HasSuffix is a wrapper for strings.HasSuffix. See HasPrefix for
// documentation.
func HasSuffix(s string) StringsConditionFunc {
	return func(a string) bool {
		return strings.HasSuffix(a, s)
	}
}
