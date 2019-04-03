// Package pie is a utility library for dealing with slices that focuses on type
// safety and performance.
//
// It can be used with the Go-style package functions:
//
//   names := []string{"Bob", "Sally", "John", "Jane"}
//   shortNames := pie.StringsOnly(names, func(s string) bool {
//   	return len(s) <= 3
//   })
//
//   // pie.Strings{"Bob"}
//
// Or, they can be chained for more complex operations:
//
//   pie.Strings{"Bob", "Sally", "John", "Jane"}.
//   	Without(pie.Prefix("J")).
//   	Transform(pie.ToUpper()).
//   	Last()
//
//   // "SALLY"
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

// StringsTransformFunc transforms a string value.
type StringsTransformFunc func(string) string

// StringsContains returns true if the string exists in the slice. The strings
// must be exactly equal (case-sensitive).
func StringsContains(ss []string, lookingFor string) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

// Contains is the chained version of StringsContains.
func (ss Strings) Contains(lookingFor string) bool {
	return StringsContains(ss, lookingFor)
}

// StringsOnly will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// StringsWithout works in the opposite way as StringsOnly.
func StringsOnly(ss []string, condition StringsConditionFunc) (ss2 []string) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Only is the chained version of StringsOnly.
func (ss Strings) Only(condition StringsConditionFunc) (ss2 Strings) {
	return StringsOnly(ss, condition)
}

// StringsWithout works the same as StringsOnly, with a negated condition. That
// is, it will return a new slice only containing the elements that returned
// false from the condition. The returned slice may contain zero elements (nil).
func StringsWithout(ss []string, condition StringsConditionFunc) (ss2 []string) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Without is the chained version of StringsWithout.
func (ss Strings) Without(condition StringsConditionFunc) (ss2 Strings) {
	return StringsWithout(ss, condition)
}

// StringsTransform will return a new slice where each element has been
// transformed. The number of element returned will always be the same as the
// input.
func StringsTransform(ss []string, fn StringsTransformFunc) (ss2 []string) {
	if ss == nil {
		return nil
	}

	ss2 = make([]string, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

// Transform is the chained version of StringsTransform.
func (ss Strings) Transform(fn StringsTransformFunc) (ss2 Strings) {
	return StringsTransform(ss, fn)
}

// StringsFirstOr returns the first element or a default value if there are no
// elements.
func StringsFirstOr(ss []string, defaultValue string) string {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// FirstOr is the chained version of StringsFirstOr.
func (ss Strings) FirstOr(defaultValue string) string {
	return StringsFirstOr(ss, defaultValue)
}

// StringsLastOr returns the last element or a default value if there are no
// elements.
func StringsLastOr(ss []string, defaultValue string) string {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// LastOr is the chained version of StringsLastOr.
func (ss Strings) LastOr(defaultValue string) string {
	return StringsLastOr(ss, defaultValue)
}

// StringsFirst returns the first element, or an empty string. Also see
// StringsFirstOr.
func StringsFirst(ss []string) string {
	return StringsFirstOr(ss, "")
}

// First is the chained version of StringsFirst.
func (ss Strings) First() string {
	return StringsFirst(ss)
}

// StringsLast returns the last element, or an empty string. Also see
// StringsLastOr.
func StringsLast(ss []string) string {
	return StringsLastOr(ss, "")
}

// Last is the chained version of StringsLast.
func (ss Strings) Last() string {
	return StringsLast(ss)
}
