package pie

import "strings"

// Prefix is a wrapper for strings.HasPrefix that can be used with any function
// that accepts as StringsConditionFunc. For example:
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
func Prefix(s string) StringsConditionFunc {
	return func(a string) bool {
		return strings.HasPrefix(a, s)
	}
}

// Suffix is a wrapper for strings.HasSuffix. See HasPrefix for documentation.
func Suffix(s string) StringsConditionFunc {
	return func(a string) bool {
		return strings.HasSuffix(a, s)
	}
}
