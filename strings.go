// Package pie is a utility library for dealing with slices that focuses only on
// type safety and performance:
//
//   shortNames := pie.Strings{"Bob", "Sally", "John", "Jane"}.
//   	Only(func(s string) bool {
//   		return len(s) <= 3
//   	})
//
// There are also helper methods for common filters, like:
//
//   namesNotStartingWithJ := pie.Strings{"Bob", "Sally", "John", "Jane"}.
//   	Without(pie.Prefix("J"))
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

// StringsApplyFunc transforms a string value.
type StringsTransformFunc func(string) string

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

// Only will return a new slice containing only the elements that return true
// from the condition. The returned slice may contain zero elements (nil).
//
// Without works in the opposite way as Only.
func (ss Strings) Only(condition StringsConditionFunc) (ss2 Strings) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Without works the same as Only, with a negated condition. That it, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss Strings) Without(condition StringsConditionFunc) (ss2 Strings) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Transform will return a new slice where each element has been transformed.
// The number of element returned will always be the same as the input.
func (ss Strings) Transform(fn StringsTransformFunc) (ss2 Strings) {
	for _, s := range ss {
		ss2 = append(ss2, fn(s))
	}

	return
}

// First returns the first element or a default value if there are no elements.
func (ss Strings) First(defaultValue string) string {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// Last returns the last element or a default value if there are no elements.
func (ss Strings) Last(defaultValue string) string {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}
