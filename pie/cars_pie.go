package pie

import (
	"encoding/json"
)

// Contains returns true if the element exists in the slice.
//
// When using slices of pointers it will only compare by address, not value.
func (ss cars) Contains(lookingFor car) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

// First returns the first element, or zero. Also see FirstOr().
func (ss cars) First() car {
	return ss.FirstOr(car{})
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss cars) FirstOr(defaultValue car) car {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss cars) JSONString() string {
	if ss == nil {
		return "[]"
	}

	// An error should not be possible.
	data, _ := json.Marshal(ss)

	return string(data)
}

// Last returns the last element, or zero. Also see LastOr().
func (ss cars) Last() car {
	return ss.LastOr(car{})
}

// LastOr returns the last element or a default value if there are no elements.
func (ss cars) LastOr(defaultValue car) car {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// Len returns the number of elements.
func (ss cars) Len() int {
	return len(ss)
}

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//   ss.Sort().Reverse()
//
func (ss cars) Reverse() cars {
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]car, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}

// Select will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// Unselect works in the opposite way as Select.
func (ss cars) Select(condition func(car) bool) (ss2 cars) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// ToStrings transforms each element to a string.
func (ss cars) ToStrings(transform func(car) string) Strings {
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
func (ss cars) Transform(fn func(car) car) (ss2 cars) {
	if ss == nil {
		return nil
	}

	ss2 = make([]car, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

// Unselect works the same as Select, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss cars) Unselect(condition func(car) bool) (ss2 cars) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}
