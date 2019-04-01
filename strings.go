// Package pie is a utility library for dealing with slices that focuses only on
// type safety and performance:
//
//   shortNames := pie.Strings{"Bob", "Sally", "John", "Jane"}.
//   	If(func(s string) bool {
//   		return len(s) <= 3
//   	})
//
// There are also helper methods for common filters, like:
//
//   namesNotStartingWithJ := pie.Strings{"Bob", "Sally", "John", "Jane"}.
//   	Unless(pie.HasPrefix("J"))
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

// StringsConditionFunc allows strings to be filtered or checked by value.
type StringsConditionFunc func(string) bool

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
//
// Unless works in the opposite way as If.
func (ss Strings) If(condition StringsConditionFunc) (ss2 Strings) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Unless works the same as If, with a negated condition. That it, it will
// return a new slice only containing the elements that returned false from the
// condition.
func (ss Strings) Unless(condition func(s string) bool) (ss2 Strings) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}
