package pie

import (
	"encoding/json"
	"sort"
)

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func (ss Float64s) All(fn func(value float64) bool) bool {
	for _, value := range ss {
		if !fn(value) {
			return false
		}
	}

	return true
}

// Any will return true if any callbacks return true. It follows the same logic
// as the any() function in Python.
//
// If the list is empty then false is always returned.
func (ss Float64s) Any(fn func(value float64) bool) bool {
	for _, value := range ss {
		if fn(value) {
			return true
		}
	}

	return false
}

// Append will return a new slice with the elements appended to the end. It is a
// wrapper for the internal append(). It is offered as a function so that it can
// more easily chained.
//
// It is acceptable to provide zero arguments.
func (ss Float64s) Append(elements ...float64) Float64s {
	return append(ss, elements...)
}

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.Float64sAreSorted.
func (ss Float64s) AreSorted() bool {
	return sort.SliceIsSorted(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}

// AreUnique will return true if the slice contains elements that are all
// different (unique) from each other.
func (ss Float64s) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
}

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss Float64s) Average() float64 {
	if l := float64(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}

// Contains returns true if the element exists in the slice.
//
// When using slices of pointers it will only compare by address, not value.
func (ss Float64s) Contains(lookingFor float64) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

// Extend will return a new slice with the slices of elements appended to the
// end.
//
// It is acceptable to provide zero arguments.
func (ss Float64s) Extend(slices ...Float64s) (ss2 Float64s) {
	ss2 = ss

	for _, slice := range slices {
		ss2 = ss2.Append(slice...)
	}

	return ss2
}

// First returns the first element, or zero. Also see FirstOr().
func (ss Float64s) First() float64 {
	return ss.FirstOr(0)
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss Float64s) FirstOr(defaultValue float64) float64 {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss Float64s) JSONString() string {
	if ss == nil {
		return "[]"
	}

	// An error should not be possible.
	data, _ := json.Marshal(ss)

	return string(data)
}

// Last returns the last element, or zero. Also see LastOr().
func (ss Float64s) Last() float64 {
	return ss.LastOr(0)
}

// LastOr returns the last element or a default value if there are no elements.
func (ss Float64s) LastOr(defaultValue float64) float64 {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// Len returns the number of elements.
func (ss Float64s) Len() int {
	return len(ss)
}

// Max is the maximum value, or zero.
func (ss Float64s) Max() (max float64) {
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

// Min is the minimum value, or zero.
func (ss Float64s) Min() (min float64) {
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

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//   ss.Sort().Reverse()
//
func (ss Float64s) Reverse() Float64s {
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]float64, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}

// Select will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// Unselect works in the opposite way as Select.
func (ss Float64s) Select(condition func(float64) bool) (ss2 Float64s) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Sort works similar to sort.Float64s(). However, unlike sort.Float64s the
// slice returned will be reallocated as to not modify the input slice.
//
// See Reverse() and AreSorted().
func (ss Float64s) Sort() Float64s {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]float64, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// Sum is the sum of all of the elements.
func (ss Float64s) Sum() (sum float64) {
	for _, s := range ss {
		sum += s
	}

	return
}

// ToStrings transforms each element to a string.
func (ss Float64s) ToStrings(transform func(float64) string) Strings {
	l := len(ss)

	// Avoid the allocation.
	if l == 0 {
		return nil
	}

	result := make(Strings, l)
	for i := 0; i < l; i++ {
		result[i] = transform(ss[i])
	}

	return result
}

// Transform will return a new slice where each element has been transformed.
// The number of element returned will always be the same as the input.
//
// Be careful when using this with slices of pointers. If you modify the input
// value it will affect the original slice. Be sure to return a new allocated
// object or deep copy the existing one.
func (ss Float64s) Transform(fn func(float64) float64) (ss2 Float64s) {
	if ss == nil {
		return nil
	}

	ss2 = make([]float64, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

// Unique returns a new slice with all of the unique values.
//
// The items will be returned in a randomized order, even with the same input.
//
// The number of items returned may be the same as the input or less. It will
// never return zero items unless then input slice has zero items.
//
// A slice with zero elements is considered to be unique.
//
// See AreUnique().
func (ss Float64s) Unique() Float64s {
	// Avoid the allocation. If there is one element or less it is already
	// unique.
	if len(ss) < 2 {
		return ss
	}

	values := map[float64]struct{}{}

	for _, value := range ss {
		values[value] = struct{}{}
	}

	var uniqueValues Float64s
	for value := range values {
		uniqueValues = append(uniqueValues, value)
	}

	return uniqueValues
}

// Unselect works the same as Select, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss Float64s) Unselect(condition func(float64) bool) (ss2 Float64s) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}
