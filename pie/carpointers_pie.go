package pie

import (
	"encoding/json"
	"github.com/elliotchance/pie/pie/util"
	"math/rand"
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

// Append will return a new slice with the elements appended to the end.
//
// It is acceptable to provide zero arguments.
func (ss carPointers) Append(elements ...*car) (result carPointers) {
	return append(append(result, ss...), elements...)
}

// Bottom will return n elements from bottom
//
// that means that elements is taken from the end of the slice
// for this [1,2,3] slice with n == 2 will be returned [3,2]
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss carPointers) Bottom(n int) (top carPointers) {
	var lastIndex = len(ss) - 1
	for i := lastIndex; i > -1 && n > 0; i-- {
		top = append(top, ss[i])
		n--
	}

	return
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

// Each is more condensed version of Transform that allows an action to happen
// on each elements and pass the original slice on.
//
//   cars.Each(func (car *Car) {
//       fmt.Printf("Car color is: %s\n", car.Color)
//   })
//
// Pie will not ensure immutability on items passed in so they can be
// manipulated, if you choose to do it this way, for example:
//
//   // Set all car colors to Red.
//   cars.Each(func (car *Car) {
//       car.Color = "Red"
//   })
//
func (ss carPointers) Each(fn func(*car)) carPointers {
	for _, s := range ss {
		fn(s)
	}

	return ss
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

// Random returns a random element by your rand.Source, or zero
func (ss carPointers) Random(source rand.Source) *car {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 1 {
		return &car{}
	}
	if n < 2 {
		return ss[0]
	}
	rnd := rand.New(source)
	i := rnd.Intn(n)
	return ss[i]
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

// Shuffle returns shuffled slice by your rand.Source
func (ss carPointers) Shuffle(source rand.Source) carPointers {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 2 {
		return ss
	}

	// go 1.10+ provides rnd.Shuffle. However, to support older versions we copy
	// the algorithm directly from the go source: src/math/rand/rand.go below,
	// with some adjustments:
	shuffled := make([]*car, n)
	copy(shuffled, ss)

	rnd := rand.New(source)

	util.Shuffle(rnd, n, func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}

// Top will return n elements from head of the slice
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss carPointers) Top(n int) (top carPointers) {
	for i := 0; i < len(ss) && n > 0; i++ {
		top = append(top, ss[i])
		n--
	}

	return
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
