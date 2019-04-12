package main

import (
	"encoding/json"
)

// The functions in this file work for all slices types.

// Contains returns true if the element exists in the slice.
//
// When using slices of pointers it will only compare by address, not value.
func (ss SliceType) Contains(lookingFor ElementType) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

// Select will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// Unselect works in the opposite way as Select.
func (ss SliceType) Select(condition func(ElementType) bool) (ss2 SliceType) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Unselect works the same as Select, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss SliceType) Unselect(condition func(ElementType) bool) (ss2 SliceType) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Transform will return a new slice where each element has been transformed.
// The number of element returned will always be the same as the input.
//
// Be careful when using this with slices of pointers. If you modify the input
// value it will affect the original slice. Be sure to return a new allocated
// object or deep copy the existing one.
func (ss SliceType) Transform(fn func(ElementType) ElementType) (ss2 SliceType) {
	if ss == nil {
		return nil
	}

	ss2 = make([]ElementType, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss SliceType) FirstOr(defaultValue ElementType) ElementType {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// LastOr returns the last element or a default value if there are no elements.
func (ss SliceType) LastOr(defaultValue ElementType) ElementType {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// First returns the first element, or zero. Also see FirstOr().
func (ss SliceType) First() ElementType {
	return ss.FirstOr(ElementZeroValue)
}

// Last returns the last element, or zero. Also see LastOr().
func (ss SliceType) Last() ElementType {
	return ss.LastOr(ElementZeroValue)
}

// Len returns the number of elements.
func (ss SliceType) Len() int {
	return len(ss)
}

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss SliceType) JSONString() string {
	if ss == nil {
		return "[]"
	}

	// An error should not be possible.
	data, _ := json.Marshal(ss)

	return string(data)
}

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//   ss.Sort().Reverse()
//
func (ss SliceType) Reverse() SliceType {
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]ElementType, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}
