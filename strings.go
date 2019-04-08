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

import (
	"encoding/json"
	"sort"
)

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

// Only will return a new slice containing only the elements that return true
// from the condition. The returned slice may contain zero elements (nil).
//
// StringsWithout works in the opposite way as StringsOnly.
func (ss Strings) Only(condition StringConditionFunc) (ss2 Strings) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Without works the same as StringsOnly, with a negated condition. That is, it
// will return a new slice only containing the elements that returned false from
// the condition. The returned slice may contain zero elements (nil).
func (ss Strings) Without(condition StringConditionFunc) (ss2 Strings) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Transform will return a new slice where each element has been transformed.
// The number of element returned will always be the same as the input.
func (ss Strings) Transform(fn StringTransformFunc) (ss2 Strings) {
	if ss == nil {
		return nil
	}

	ss2 = make([]string, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss Strings) FirstOr(defaultValue string) string {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// LastOr returns the last element or a default value if there are no elements.
func (ss Strings) LastOr(defaultValue string) string {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// First returns the first element, or an empty string. Also see FirstOr().
func (ss Strings) First() string {
	return ss.FirstOr("")
}

// Last returns the last element, or an empty string. Also see LastOr().
func (ss Strings) Last() string {
	return ss.LastOr("")
}

// Len returns the number of elements.
func (ss Strings) Len() int {
	return len(ss)
}

// Min is the minimum value, or an empty string.
func (ss Strings) Min() (min string) {
	if len(ss) == 0 {
		return
	}

	min = ss[0]
	for _, s := range ss {
		if s < min {
			min = s
		}
	}

	return
}

// Max is the maximum value, or en empty string.
func (ss Strings) Max() (max string) {
	if len(ss) == 0 {
		return
	}

	max = ss[0]
	for _, s := range ss {
		if s > max {
			max = s
		}
	}

	return
}

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss Strings) JSONString() string {
	if ss == nil {
		return "[]"
	}

	// An error should not be possible.
	data, _ := json.Marshal(ss)

	return string(data)
}

// Sort works similar to sort.Strings(). However, unlike sort.Strings the
// slice returned will be reallocated as to not modify the input slice.
//
// See Reverse() and AreSorted().
func (ss Strings) Sort() Strings {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]string, len(ss))
	copy(sorted, ss)
	sort.Strings(sorted)

	return sorted
}

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//   ss.Sort().Reverse()
//
func (ss Strings) Reverse() Strings {
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]string, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.StringsAreSorted.
func (ss Strings) AreSorted() bool {
	return sort.StringsAreSorted(ss)
}

// Unique returns a slice that only contains unique elements.
//
// The unique algorithm requires that all items be sorted first so the returned
// slice will not be in the same order as the input slice.
func (ss Strings) Unique() Strings {
	// Avoid the allocation. If there is one element or less it is already
	// unique.
	if len(ss) < 2 {
		return ss
	}

	sorted := ss.Sort()
	unique := Strings{sorted[0]}
	i := 0
	for _, value := range sorted {
		if unique[i] != value {
			unique = append(unique, value)
			i++
		}
	}

	return unique
}

// AreUnique returns true if all of the elements are unique.
func (ss Strings) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
}
