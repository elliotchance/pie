package pie

import (
	"encoding/json"
)

// The functions in this file work for all slices types.

// Contains returns true if the element exists in the slice.
func (ss carPointers) Contains(lookingFor *car) bool {
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
// carPointersWithout works in the opposite way as carPointersOnly.
func (ss carPointers) Only(condition func(*car) bool) (ss2 carPointers) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Without works the same as Only, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss carPointers) Without(condition func(*car) bool) (ss2 carPointers) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Transform will return a new slice where each element has been transformed.
// The number of element returned will always be the same as the input.
func (ss carPointers) Transform(fn func(*car) *car) (ss2 carPointers) {
	if ss == nil {
		return nil
	}

	ss2 = make([]*car, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss carPointers) FirstOr(defaultValue *car) *car {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// LastOr returns the last element or a default value if there are no elements.
func (ss carPointers) LastOr(defaultValue *car) *car {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// First returns the first element, or zero. Also see FirstOr().
func (ss carPointers) First() *car {
	return ss.FirstOr(&car{})
}

// Last returns the last element, or zero. Also see LastOr().
func (ss carPointers) Last() *car {
	return ss.LastOr(&car{})
}

// Len returns the number of elements.
func (ss carPointers) Len() int {
	return len(ss)
}

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss carPointers) JSONString() string {
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
func (ss carPointers) Reverse() carPointers {
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]*car, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}
