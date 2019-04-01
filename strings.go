// Package pie is a utility library for dealing with slices that focuses only on
// type safety and performance:
//
//   namesStartingWithJ := pie.Strings{"Bob", "Sally", "John", "Jane"}.
//   	If(func(s string) bool {
//   		return s[0] == 'J'
//   	})
//
package pie

// Strings is an alias for a string slice.
//
// You can create a Strings directly:
//
//   pie.Strings{"foo", "bar"}
//
// Or, cast an existing string slice:
//
//   ss := []string{"foo", "bar"}
//   pie.Strings(ss)
//
type Strings []string

// Contains returns true if the string exists in the slice. The strings must be
// exactly equal (case-sensitive).
func (ss Strings) Contains(lookingFor string) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

// If will return a new slice containing only the elements that return true from
// the condition. The returned slice may contain zero elements (nil).
func (ss Strings) If(condition func(s string) bool) (ss2 Strings) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}
