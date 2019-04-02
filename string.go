package pie

import "strings"

// HasPrefix is a wrapper for strings.HasPrefix that can be used with any
// function that accepts a StringsConditionFunc. For example:
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

// HasSuffix is a wrapper for strings.HasSuffix. See HasPrefix for
// documentation.
func Suffix(s string) StringsConditionFunc {
	return func(a string) bool {
		return strings.HasSuffix(a, s)
	}
}

// ToUpper is a wrapper for strings.ToUpper and can be used with any function
// that accepts StringsTransformFunc:
//
//   pie.Strings{"Bob", "Sally", "John", "Jane"}.
//   	Transform(pie.ToUpper)
//
func ToUpper() StringsTransformFunc {
	return func (s string) string {
		return strings.ToUpper(s)
	}
}

// ToLower is a wrapper for strings.ToLower and can be used with any function
// that accepts StringsTransformFunc:
//
//   pie.Strings{"Bob", "Sally", "John", "Jane"}.
//   	Transform(pie.ToLower)
//
func ToLower() StringsTransformFunc {
	return func (s string) string {
		return strings.ToLower(s)
	}
}
