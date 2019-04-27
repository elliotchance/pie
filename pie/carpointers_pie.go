package pie

import (
	"encoding/json"
	"math/rand"
	"time"
)

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func (ss carPointers) All(fn func(value *car) bool) bool {
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
func (ss carPointers) Any(fn func(value *car) bool) bool {
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
func (ss carPointers) Append(elements ...*car) carPointers {
	return append(ss, elements...)
}

// Contains returns true if the element exists in the slice.
//
// When using slices of pointers it will only compare by address, not value.
func (ss carPointers) Contains(lookingFor *car) bool {
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
func (ss carPointers) Extend(slices ...carPointers) (ss2 carPointers) {
	ss2 = ss

	for _, slice := range slices {
		ss2 = ss2.Append(slice...)
	}

	return ss2
}

// First returns the first element, or zero. Also see FirstOr().
func (ss carPointers) First() *car {
	return ss.FirstOr(&car{})
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss carPointers) FirstOr(defaultValue *car) *car {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
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

// Last returns the last element, or zero. Also see LastOr().
func (ss carPointers) Last() *car {
	return ss.LastOr(&car{})
}

// LastOr returns the last element or a default value if there are no elements.
func (ss carPointers) LastOr(defaultValue *car) *car {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// Len returns the number of elements.
func (ss carPointers) Len() int {
	return len(ss)
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

// Select will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// Unselect works in the opposite way as Select.
func (ss carPointers) Select(condition func(*car) bool) (ss2 carPointers) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Max is the maximum value, or zero.
func (ss carPointers) Shuffle() carPointers {
	if len(ss) < 2 {
		return ss
	}

	shuffled := make([]*car, len(ss))
	copy(shuffled, ss)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}

// ToStrings transforms each element to a string.
func (ss carPointers) ToStrings(transform func(*car) string) Strings {
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

// Unselect works the same as Select, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss carPointers) Unselect(condition func(*car) bool) (ss2 carPointers) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}
