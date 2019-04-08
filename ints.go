package pie

import (
	"encoding/json"
	"sort"
)

// Ints is an alias for an int slice.
//
// You can create an Ints directly:
//
//   pie.Ints{1, 2, 3}
//
// Or, cast an existing int slice:
//
//   ss := []int{1, 2, 3}
//   pie.Ints(ss)
//
type Ints []int

// Contains returns true if the int exists in the slice.
func (ss Ints) Contains(lookingFor int) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

// Only will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// IntsWithout works in the opposite way as IntsOnly.
func (ss Ints) Only(condition IntConditionFunc) (ss2 Ints) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Without works the same as IntsOnly, with a negated condition. That is, it
// will return a new slice only containing the elements that returned false from
// the condition. The returned slice may contain zero elements (nil).
func (ss Ints) Without(condition IntConditionFunc) (ss2 Ints) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Transform will return a new slice where each element has been transformed.
// The number of element returned will always be the same as the input.
func (ss Ints) Transform(fn IntTransformFunc) (ss2 Ints) {
	if ss == nil {
		return nil
	}

	ss2 = make([]int, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss Ints) FirstOr(defaultValue int) int {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// LastOr returns the last element or a default value if there are no elements.
func (ss Ints) LastOr(defaultValue int) int {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// First returns the first element, or zero. Also see FirstOr().
func (ss Ints) First() int {
	return ss.FirstOr(0)
}

// Last returns the last element, or zero. Also see LastOr().
func (ss Ints) Last() int {
	return ss.LastOr(0)
}

// Sum is the sum of all of the elements.
func (ss Ints) Sum() (sum int) {
	for _, s := range ss {
		sum += s
	}

	return
}

// Len returns the number of elements.
func (ss Ints) Len() int {
	return len(ss)
}

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss Ints) Average() float64 {
	if l := float64(len(ss)); l > 0 {
		return float64(ss.Sum()) / l
	}

	return 0
}

// Min is the minimum value, or zero.
func (ss Ints) Min() (min int) {
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

// Max is the maximum value, or zero.
func (ss Ints) Max() (max int) {
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
func (ss Ints) JSONString() string {
	if ss == nil {
		return "[]"
	}

	// An error should not be possible.
	data, _ := json.Marshal(ss)

	return string(data)
}

// Sort works similar to sort.Ints(). However, unlike sort.Ints the slice
// returned will be reallocated as to not modify the input slice.
//
// See Reverse() and AreSorted().
func (ss Ints) Sort() Ints {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]int, len(ss))
	copy(sorted, ss)
	sort.Ints(sorted)

	return sorted
}

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//   ss.Sort().Reverse()
//
func (ss Ints) Reverse() Ints {
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]int, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.IntsAreSorted.
func (ss Ints) AreSorted() bool {
	return sort.IntsAreSorted(ss)
}
